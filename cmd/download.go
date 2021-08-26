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

	"github.com/IntegrationWorks/kobi/internal"
	"github.com/spf13/cobra"
)

var (
	downloadCmdService string
	downloadCmdOutput  string
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download [service]",
	Short: "Download the BIAN API Definition for the identified servce.",
	Long: `Download the BIAN Swagger Specification (.json) file from the public
BIAN repository of semantic APIs. By default, the file will be saved
in the current working directory, this can be overridden by the -o 
or --output flag to reference a filepath to save the content to.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 || args[0] == "" {
			return fmt.Errorf("missing argument `service`")
		}
		downloadCmdService = args[0]
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("download called for service %s\n", downloadCmdService)
		return internal.DownloadFile(downloadCmdService, downloadCmdOutput)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringVarP(&downloadCmdOutput, "output", "o", "", "Output file to save BIAN API Specification to.")
}
