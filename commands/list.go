package commands

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/bradfitz/slice"
	"github.com/fatih/color"
	apischema "github.com/giantswarm/api-schema"
	"github.com/giantswarm/gsclientgen"
	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"

	"github.com/giantswarm/gsctl/config"
	"github.com/giantswarm/gsctl/util"
)

var (
	// ListCommand is the command to list things
	ListCommand = &cobra.Command{
		Use:   "list",
		Short: "List things, like organizations, clusters, key-pairs",
		Long:  `Prints a list of the things you have access to`,
	}

	// ListOrgsCommand performs the "list organizations" function
	ListOrgsCommand = &cobra.Command{
		Use:     "organizations",
		Short:   "List organizations",
		Long:    `Prints a list of the organizations you are a member of`,
		PreRunE: checkListOrgs,
		Run:     listOrgs,
	}

	// ListClustersCommand performs the "list clusters" function
	ListClustersCommand = &cobra.Command{
		Use:     "clusters",
		Short:   "List clusters",
		Long:    `Prints a list of all clusters you have access to`,
		PreRunE: checkListClusters,
		Run:     listClusters,
	}

	// ListKeypairsCommand performs the "list keypairs" function
	ListKeypairsCommand = &cobra.Command{
		Use:     "keypairs",
		Short:   "List key-pairs for a cluster",
		Long:    `Prints a list of key-pairs for a cluster`,
		PreRunE: checkListKeypairs,
		Run:     listKeypairs,
	}
)

func init() {
	ListKeypairsCommand.Flags().StringVarP(&cmdClusterID, "cluster", "c", "", "ID of the cluster to list key-pairs for")
	// subcommands
	ListCommand.AddCommand(ListOrgsCommand, ListClustersCommand, ListKeypairsCommand)

	RootCommand.AddCommand(ListCommand)
}

func checkListOrgs(cmd *cobra.Command, args []string) error {
	if config.Config.Token == "" && cmdToken == "" {
		return errors.New("You are not logged in.\nUse '" + config.ProgramName + " login' to login or '--auth-token' to pass a valid auth token.")
	}
	return nil
}

// list organizations the user is member of
func listOrgs(cmd *cobra.Command, args []string) {
	client := gsclientgen.NewDefaultApi()

	// if token is set via flags, we unauthenticate using this token
	authHeader := "giantswarm " + config.Config.Token
	if cmdToken != "" {
		authHeader = "giantswarm " + cmdToken
	}

	orgsResponse, apiResponse, err := client.GetUserOrganizations(authHeader)
	if err != nil {
		fmt.Println("Error details:")
		log.Fatal(err)
	}

	if orgsResponse.StatusCode == apischema.STATUS_CODE_DATA {
		var organizations = orgsResponse.Data
		if len(organizations) == 0 {
			fmt.Println(color.YellowString("No organizations available"))
		} else {
			sort.Strings(organizations)
			fmt.Println(color.YellowString("Organization"))
			for _, orgName := range organizations {
				fmt.Println(color.CyanString(orgName))
			}
		}
	} else {
		fmt.Printf("Unhandled response code: %v", orgsResponse.StatusCode)
		fmt.Printf("Status text: %v", orgsResponse.StatusText)
		fmt.Printf("apiResponse: %s\n", apiResponse)
	}
}

func checkListClusters(cmd *cobra.Command, args []string) error {
	if config.Config.Token == "" {
		return errors.New("You are not logged in. Use '" + config.ProgramName + " login' to log in.")
	}
	return nil
}

// list all clusters the user has access to
func listClusters(cmd *cobra.Command, args []string) {
	client := gsclientgen.NewDefaultApi()
	authHeader := "giantswarm " + config.Config.Token
	orgsResponse, apiResponse, err := client.GetUserOrganizations(authHeader)
	if err != nil {
		log.Fatal(err)
	}
	if orgsResponse.StatusCode == apischema.STATUS_CODE_DATA {
		var organizations = orgsResponse.Data
		if len(organizations) == 0 {
			fmt.Println(color.YellowString("No organizations available"))
		} else {
			sort.Strings(organizations)
			output := []string{color.YellowString("Id") + "|" + color.YellowString("Name") + "|" + color.YellowString("Created") + "|" + color.YellowString("Organization")}
			for _, orgName := range organizations {
				clustersResponse, _, err := client.GetOrganizationClusters(authHeader, orgName)
				if err != nil {
					log.Fatal(err)
				}
				for _, cluster := range clustersResponse.Data.Clusters {
					created := util.ShortDate(util.ParseDate(cluster.CreateDate))
					output = append(output,
						color.CyanString(cluster.Id)+"|"+
							color.CyanString(cluster.Name)+"|"+
							color.CyanString(created)+"|"+
							color.CyanString(orgName))
				}
			}
			fmt.Println(columnize.SimpleFormat(output))
		}
	} else {
		fmt.Printf("Unhandled response code: %v", orgsResponse.StatusCode)
		fmt.Printf("Status text: %v", orgsResponse.StatusText)
		fmt.Printf("apiResponse: %s\n", apiResponse)
	}
}

func checkListKeypairs(cmd *cobra.Command, args []string) error {
	if config.Config.Token == "" {
		return errors.New("You are not logged in. Use '" + config.ProgramName + " login' to log in.")
	}
	if cmdClusterID == "" {
		// use default cluster if possible
		clusterID, _ := config.GetDefaultCluster()
		if clusterID != "" {
			cmdClusterID = clusterID
		} else {
			return errors.New("No cluster given. Please use the -c/--cluster flag to set a cluster ID.")
		}
	}
	return nil
}

func listKeypairs(cmd *cobra.Command, args []string) {
	client := gsclientgen.NewDefaultApi()
	authHeader := "giantswarm " + config.Config.Token
	keypairsResponse, _, err := client.GetKeyPairs(authHeader, cmdClusterID)
	if err != nil {
		log.Fatal(err)
	}
	if keypairsResponse.StatusCode == apischema.STATUS_CODE_DATA {
		// sort result
		slice.Sort(keypairsResponse.Data.KeyPairs[:], func(i, j int) bool {
			return keypairsResponse.Data.KeyPairs[i].CreateDate < keypairsResponse.Data.KeyPairs[j].CreateDate
		})

		// create output
		output := []string{color.YellowString("Created") + "|" + color.YellowString("Expires") + "|" + color.YellowString("Id") + "|" + color.YellowString("Description")}
		for _, keypair := range keypairsResponse.Data.KeyPairs {
			created := util.ShortDate(util.ParseDate(keypair.CreateDate))
			expires := util.ParseDate(keypair.CreateDate).Add(time.Duration(keypair.TtlHours) * time.Hour)

			// skip if expired
			output = append(output, color.CyanString(created)+"|"+
				color.CyanString(util.ShortDate(expires))+"|"+
				color.CyanString(util.Truncate(util.CleanKeypairID(keypair.Id), 10))+"|"+
				color.CyanString(keypair.Description))
		}
		fmt.Println(columnize.SimpleFormat(output))

	} else {
		fmt.Printf("Unhandled response code: %v", keypairsResponse.StatusCode)
		fmt.Printf("Status text: %v", keypairsResponse.StatusText)
	}
}