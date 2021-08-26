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
	"github.com/IntegrationWorks/kobi/internal"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a Kong Developer Portal project template with a BIAN workspace",
	Long: `Create a project framework using the Kong Developer Portal template. Create a
BIAN workspace directory to hold the BIAN API Specifications and thematic files.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return internal.InitialiseWorkspace()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
