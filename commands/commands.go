package commands

// This file defines some variables to be available in all commands

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/giantswarm/microerror"

	"github.com/giantswarm/gsctl/client"
	"github.com/giantswarm/gsctl/client/clienterror"
	"github.com/giantswarm/gsctl/oidc"
)

var (
	// API endpoint flag
	cmdAPIEndpoint string

	// token flag
	cmdToken string

	// configuration path to use temporarily
	cmdConfigDirPath string

	// verbose flag
	cmdVerbose bool

	// certificate-organizations flag
	cmdCertificateOrganizations string

	// cluster ID flag
	cmdClusterID string

	// cn-prefix flag
	cmdCNPrefix string

	// description flag
	cmdDescription string

	// TTL (time to live) flag
	cmdTTL string

	// force flag. if set, no prompt should be displayed.
	cmdForce bool

	// full flag. if set, output must not be truncated.
	cmdFull bool

	// organization ID as passed by the user as a flag
	cmdOrganizationID string

	// cmdRelease sets a release to use, provided as a command line flag
	cmdRelease string

	// number of CPUs per worker as required via flag on execution
	cmdWorkerNumCPUs int

	// RAM size per worker node in GB per worker as required via flag on execution
	cmdWorkerMemorySizeGB float32

	// Local storage per worker node in GB per worker as required via flag on execution
	cmdWorkerStorageSizeGB float32
)

type cpuDefinition struct {
	Cores int `yaml:"cores,omitempty"`
}

type memoryDefinition struct {
	SizeGB float32 `yaml:"size_gb,omitempty"`
}

type storageDefinition struct {
	SizeGB float32 `yaml:"size_gb,omitempty"`
}

type awsSpecificDefinition struct {
	InstanceType string `yaml:"instance_type,omitempty"`
}

type azureSpecificDefinition struct {
	VMSize string `yaml:"vm_size,omitempty"`
}

type nodeDefinition struct {
	Memory  memoryDefinition        `yaml:"memory,omitempty"`
	CPU     cpuDefinition           `yaml:"cpu,omitempty"`
	Storage storageDefinition       `yaml:"storage,omitempty"`
	Labels  map[string]string       `yaml:"labels,omitempty"`
	AWS     awsSpecificDefinition   `yaml:"aws,omitempty"`
	Azure   azureSpecificDefinition `yaml:"azure,omitempty"`
}

type clusterDefinition struct {
	Name              string            `yaml:"name,omitempty"`
	Owner             string            `yaml:"owner,omitempty"`
	ReleaseVersion    string            `yaml:"release_version,omitempty"`
	AvailabilityZones int               `yaml:"availability_zones,omitempty"`
	Scaling           scalingDefinition `yaml:"scaling,omitempty"`
	Workers           []nodeDefinition  `yaml:"workers,omitempty"`
}

type scalingDefinition struct {
	Min int64 `yaml:"min,omitempty"`
	Max int64 `yaml:"max,omitempty"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// askForConfirmation asks the user for confirmation. A user must type in "yes" or "no" and
// then press enter. It has fuzzy matching, so "y", "Y", "yes", "YES", and "Yes" all count as
// confirmations. If the input is not recognized, it will ask again. The function does not return
// until it gets a valid response from the user.
func askForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", color.YellowString(s))

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}

// handleCommonErrors is a common function to handle certain errors happening in
// more than one command. If the error given is handled by the function, it
// prints according text for the end user and exits the process.
// If the error is not recognized, we simply return.
//
func handleCommonErrors(err error) {

	var headline = ""
	var subtext = ""

	// V2 client error handling
	if convertedErr, ok := microerror.Cause(err).(*clienterror.APIError); ok {
		headline = convertedErr.ErrorMessage
		subtext = convertedErr.ErrorDetails
	} else if convertedErr, ok := err.(*clienterror.APIError); ok {
		headline = convertedErr.ErrorMessage
		subtext = convertedErr.ErrorDetails
	} else {
		// legacy client error handling
		switch {
		case client.IsEndpointNotSpecifiedError(err):
			headline = "No endpoint has been specified."
			subtext = "Please use the '-e|--endpoint' flag."
		case oidc.IsAuthorizationError(err):
			headline = "Unauthorized"
			subtext = "Something went wrong during a OIDC operation: " + err.Error() + "\n"
			subtext += "Please try logging in again."
		case oidc.IsRefreshError(err):
			headline = "Unable to refresh your SSO token."
			subtext = err.Error() + "\n"
			subtext += "Please try loging in again using: gsctl login --sso"
		case IsNotLoggedInError(err):
			headline = "You are not logged in."
			subtext = "Use 'gsctl login' to login or '--auth-token' to pass a valid auth token."
		case IsAccessForbiddenError(err):
			// TODO: remove once the legacy client is no longer used
			headline = "Access Forbidden"
			subtext = "The client has been denied access to the API endpoint with an HTTP status of 403.\n"
			subtext += "Please make sure that you are in the right network or VPN. Once that is verified,\n"
			subtext += "check back with Giant Swarm support that your network is permitted access."
		case IsEmptyPasswordError(err):
			headline = "Empty password submitted"
			subtext = "The API server complains about the password provided."
			subtext += " Please make sure to provide a string with more than white space characters."
		case IsClusterIDMissingError(err):
			headline = "No cluster ID specified."
			subtext = "Please specify a cluster ID. Use --help for details."
		case IsCouldNotCreateClientError(err):
			headline = "Failed to create API client."
			subtext = "Details: " + err.Error()
		case IsNotAuthorizedError(err):
			// TODO: remove once the legacy client is no longer used
			headline = "You are not authorized for this action."
			subtext = "Please check whether you are logged in with the right credentials using 'gsctl info'."
		case IsInternalServerError(err):
			headline = "An internal error occurred."
			subtext = "Please try again in a few minutes. If that does not success, please inform the Giant Swarm support team."
		case IsNoResponseError(err):
			headline = "The API didn't send a response."
			subtext = "Please check your connection using 'gsctl ping'. If your connection is fine,\n"
			subtext += "please try again in a few moments."
		case IsUnknownError(err):
			headline = "An error occurred."
			subtext = "Please notify the Giant Swarm support team, or try the command again in a few moments.\n"
			subtext += fmt.Sprintf("Details: %s", err.Error())
		}

	}

	if headline == "" {
		return
	}

	fmt.Println(color.RedString(headline))
	if subtext != "" {
		fmt.Println(subtext)
	}
	os.Exit(1)
}
