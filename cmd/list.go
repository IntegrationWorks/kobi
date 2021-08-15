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

	"github.com/IntegrationWorks/bong/internal"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available BIAN service domains",
	Long: `Returns a list of available BIAN service domains which can be
subsequently published to a Kong API Manager or Kong Developer 
Portal using bong publish `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("-- Listing all available BIAN service domains --")
		internal.ListServices()
		fmt.Println("------------------------------------------------")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
