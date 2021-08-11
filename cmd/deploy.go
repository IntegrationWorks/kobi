/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/IntegrationWorks/bong/internal"
	"github.com/spf13/cobra"
)

var (
	deployCmdWorkspace string
	deployCmdService   string
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy [service]",
	Short: "Deploy the BIAN API Definition to a Kong Developer Portal.",
	Long: `Download the BIAN Swagger Specification (.json) file from the public
BIAN repository of semantic APIs and deploy it to the Developer Portal
of an existing Kong installation. 

By default, Bong will look for the Kong admin API at http://localhost:8001
with no authentication configured. This behaviour can be overridden by setting
the environment variables "BONG_KONG_ADDR" and BONG_KONG_TOKEN".

BONG_KONG_ADDR: full base URL of Kong Admin API. E.g. https://my-kong-host:8444
BONG_KONG_TOKEN: RBAC token value for configured Kong user with write permissions
to the Portal Files API.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 || args[0] == "" {
			return fmt.Errorf("missing argument `service`")
		}
		deployCmdService = args[0]
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		err := internal.DownloadFile(deployCmdService, "")
		if err != nil {
			return err
		}
		defer os.Remove(deployCmdService + ".json")
		return internal.DeploySpecToPortal(deployCmdWorkspace, deployCmdService)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.Flags().StringVarP(&deployCmdWorkspace, "workspace", "w", "", "Kong Workspace to deploy specification to.")

}
