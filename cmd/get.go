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

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var hidePassword bool
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get returns the password stored for a given account",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Verify that a password is set for the vault
		// and a default vault exists
		initVault()

		if len(args) < 1 {
			fmt.Println("😐 No user specified")
			return
		}

		name := args[0]
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

		if _, ok := vault[name]; !ok {
			fmt.Printf("😐 No user %s in the vault", name)
			return
		}
		// copy password to clipboard
		clipboard.WriteAll(vault[name])
		if hidePassword {
			fmt.Print(vault[name])
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolVarP(&hidePassword, "verbose", "v", false, "print password to the command line")
}
