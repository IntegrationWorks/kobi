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

	"github.com/IntegrationWorks/kobi/internal"
	"github.com/spf13/cobra"
)

var (
	publishCmdWorkspace string
	publishCmdService   string
)

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish [service]",
	Short: "publish the BIAN API Definition to a Kong Developer Portal.",
	Long: `Download the BIAN Swagger Specification (.json) file from the public
BIAN repository of semantic APIs and publish it to the Developer Portal
of an existing Kong installation. 

By default, kobi will look for the Kong admin API at http://localhost:8001
with no authentication configured. This behaviour can be overridden by setting
the environment variables "kobi_KONG_ADDR" and kobi_KONG_TOKEN".

kobi_KONG_ADDR: full base URL of Kong Admin API. E.g. https://my-kong-host:8444
kobi_KONG_TOKEN: RBAC token value for configured Kong user with write permissions
to the Portal Files API.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 || args[0] == "" {
			return fmt.Errorf("missing argument `service`")
		}
		publishCmdService = args[0]
		return internal.ValidateBianVersionAndAPIType(bianVersion, apiType)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		err := internal.DownloadFile(publishCmdService, "", bianVersion, apiType)
		if err != nil {
			return err
		}
		defer os.Remove(publishCmdService + ".json")
		return internal.PublishSpecToPortal(publishCmdWorkspace, publishCmdService, bianVersion)
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
	publishCmd.Flags().StringVarP(&publishCmdWorkspace, "workspace", "w", "", "Kong Workspace to publish specification to.")
}
