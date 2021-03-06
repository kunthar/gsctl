package commands

import (
	"fmt"
	"net/http"
	"os"

	"github.com/giantswarm/microerror"

	"github.com/fatih/color"
	"github.com/giantswarm/gsctl/client/clienterror"
	"github.com/giantswarm/gsctl/config"
	"github.com/spf13/cobra"
)

type deleteClusterArguments struct {
	// API endpoint
	apiEndpoint string
	// cluster ID to delete
	clusterID string
	// cluster ID passed via -c/--cluster argument
	legacyClusterID string
	// don't prompt
	force bool
	// auth scheme
	scheme string
	// auth token
	token string
	// verbosity
	verbose bool
}

func defaultDeleteClusterArguments(positionalArgs []string) deleteClusterArguments {
	endpoint := config.Config.ChooseEndpoint(cmdAPIEndpoint)
	token := config.Config.ChooseToken(endpoint, cmdToken)
	scheme := config.Config.ChooseScheme(endpoint, cmdToken)

	clusterID := ""
	if len(positionalArgs) > 0 {
		clusterID = positionalArgs[0]
	}

	return deleteClusterArguments{
		apiEndpoint:     endpoint,
		clusterID:       clusterID,
		force:           cmdForce,
		legacyClusterID: cmdClusterID,
		scheme:          scheme,
		token:           token,
		verbose:         cmdVerbose,
	}
}

const (
	deleteClusterActivityName = "delete-cluster"
)

var (
	// DeleteClusterCommand performs the "delete cluster" function
	DeleteClusterCommand = &cobra.Command{
		Use:   "cluster",
		Short: "Delete cluster",
		Long: `Deletes a Kubernetes cluster.

Caution: This will terminate all workloads on the cluster. Data stored on the
worker nodes will be lost. There is no way to undo this.

Example:

	gsctl delete cluster c7t2o`,
		PreRun: deleteClusterValidationOutput,
		Run:    deleteClusterExecutionOutput,
	}
)

func init() {
	DeleteClusterCommand.Flags().StringVarP(&cmdClusterID, "cluster", "c", "", "ID of the cluster to delete")
	DeleteClusterCommand.Flags().BoolVarP(&cmdForce, "force", "", false, "If set, no interactive confirmation will be required (risky!).")

	DeleteClusterCommand.Flags().MarkDeprecated("cluster", "You no longer need to pass the cluster ID with -c/--cluster. Use --help for details.")

	DeleteCommand.AddCommand(DeleteClusterCommand)
}

// deleteClusterValidationOutput runs our pre-checks.
// If errors occur, error info is printed to STDOUT/STDERR
// and the program will exit with non-zero exit codes.
func deleteClusterValidationOutput(cmd *cobra.Command, args []string) {
	dca := defaultDeleteClusterArguments(args)

	err := validateDeleteClusterPreConditions(dca)
	if err != nil {
		handleCommonErrors(err)

		var headline = ""
		var subtext = ""

		switch {
		case IsConflictingFlagsError(err):
			headline = "Conflicting flags/arguments"
			subtext = "Please specify the cluster to be used as a positional argument, avoid -c/--cluster."
			subtext += "See --help for details."
		case IsClusterIDMissingError(err):
			headline = "No cluster ID specified"
			subtext = "See --help for usage details."
		case IsCouldNotDeleteClusterError(err):
			headline = "The cluster could not be deleted."
			subtext = "You might try again in a few moments. If that doesn't work, please contact the Giant Swarm support team."
			subtext += " Sorry for the inconvenience!"
		default:
			headline = err.Error()
		}

		// print output
		fmt.Println(color.RedString(headline))
		if subtext != "" {
			fmt.Println(subtext)
		}
		os.Exit(1)
	}
}

// validateDeleteClusterPreConditions checks preconditions and returns an error in case
func validateDeleteClusterPreConditions(args deleteClusterArguments) error {
	if args.clusterID == "" && args.legacyClusterID == "" {
		return microerror.Mask(clusterIDMissingError)
	}
	if args.clusterID != "" && args.legacyClusterID != "" {
		return microerror.Mask(conflictingFlagsError)
	}
	if config.Config.Token == "" && args.token == "" {
		return microerror.Mask(notLoggedInError)
	}
	return nil
}

// interprets arguments/flags, eventually submits delete request
func deleteClusterExecutionOutput(cmd *cobra.Command, args []string) {
	dca := defaultDeleteClusterArguments(args)
	deleted, err := deleteCluster(dca)
	if err != nil {
		handleCommonErrors(err)

		var headline = ""
		var subtext = ""

		switch {
		case IsClusterNotFoundError(err):
			headline = "Cluster not found"
			subtext = "The cluster you tried to delete doesn't seem to exist. Check 'gsctl list clusters' to make sure."
		default:
			headline = err.Error()
		}

		fmt.Println(color.RedString(headline))
		if subtext != "" {
			fmt.Println(subtext)
		}
		os.Exit(1)
	}

	// non-error output
	if deleted {
		clusterID := dca.legacyClusterID
		if dca.clusterID != "" {
			clusterID = dca.clusterID
		}
		fmt.Println(color.GreenString("The cluster with ID '%s' will be deleted as soon as all workloads are terminated.", clusterID))
	} else {
		if dca.verbose {
			fmt.Println(color.GreenString("Aborted."))
		}
	}
}

// deleteCluster performs the cluster deletion API call
//
// The returned tuple contains:
// - bool: true if cluster will reall ybe deleted, false otherwise
// - error: The error that has occurred (or nil)
//
func deleteCluster(args deleteClusterArguments) (bool, error) {
	// Accept legacy cluster ID for a while, but real one takes precedence.
	clusterID := args.legacyClusterID
	if args.clusterID != "" {
		clusterID = args.clusterID
	}

	// confirmation
	if !args.force {
		confirmed := askForConfirmation("Do you really want to delete cluster '" + clusterID + "'?")
		if !confirmed {
			return false, nil
		}
	}

	auxParams := ClientV2.DefaultAuxiliaryParams()
	auxParams.ActivityName = deleteClusterActivityName

	// perform API call
	_, err := ClientV2.DeleteCluster(clusterID, auxParams)
	if err != nil {
		// create specific error types for cases we care about
		if clientErr, ok := err.(*clienterror.APIError); ok {
			if clientErr.HTTPStatusCode == http.StatusForbidden {
				return false, microerror.Mask(accessForbiddenError)
			} else if clientErr.HTTPStatusCode == http.StatusNotFound {
				return false, microerror.Mask(clusterNotFoundError)
			}
		}

		return false, microerror.Maskf(couldNotDeleteClusterError, err.Error())
	}

	return true, nil
}
