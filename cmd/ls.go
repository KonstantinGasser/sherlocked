/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "displays all stored accounts",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Verify that a password is set for the vault
		// and a default vault exists
		initVault()

		// get vault password
		password, err := clIO.Password()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// get encryted vault
		encryted, err := PassManager.Read()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// decrypt vault
		vault, err := PassManager.Decrypt(password, encryted)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Stored keys:")
		for key := range vault {
			fmt.Printf("	🗝 %s\n", key)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
