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
		if err := runGet(args); err != nil {
			fmt.Println(err.Error())
			return
		}
	},
}

// run func holds the logic for the password command
// is a separated function to test the code proper
func runGet(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("😐 No user specified")
	}

	name := args[0]
	// get vault password
	password, err := clIO.Password()
	if err != nil {
		return err
	}

	// get encryted vault
	encryted, err := PassManager.Read()
	if err != nil {
		return err
	}
	// decrypt vault
	vault, err := PassManager.Decrypt(password, encryted)
	if err != nil {
		return err
	}

	if _, ok := vault[name]; !ok {
		return fmt.Errorf("😐 No user %s in the vault", name)
	}
	// copy password to clipboard
	clipboard.WriteAll(vault[name])
	if hidePassword {
		fmt.Print(vault[name])
	}
	return nil
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolVarP(&hidePassword, "verbose", "v", false, "print password to the command line")
}
