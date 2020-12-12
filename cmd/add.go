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

var override bool

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new username and password to the key vault",
	Long:  `with the add command you can add a new password mapped to a username or account name`,
	Run: func(cmd *cobra.Command, args []string) {
		// Verify that a password is set for the vault
		// and a default vault exists
		initVault()

		if err := runAdd(args); err != nil {
			fmt.Println(err.Error())
			return
		}
	},
}

// run func holds the logic for the password command
// is a separated function to test the code proper
func runAdd(args []string) error {
	password, err := clIO.Password()
	if err != nil {
		return err
	}

	uname, pass, err := clIO.Credentials()
	if err != nil {
		return err
	}

	// get encryted vault
	fileContent, err := PassManager.Read()
	if err != nil {
		return err
	}
	// decrypt vault
	vault, err := PassManager.Decrypt(password, fileContent)
	if err != nil {
		return err
	}

	// check if user exists in vault with same key
	if _, ok := vault[uname]; ok && !override {
		return fmt.Errorf("🤔 User %s already stroed use add --override (this option is inreversable) or del -user\n", uname)
	}

	// add new account to vault
	vault[uname] = pass

	// write changed vault
	b, err := PassManager.Serialize(vault)
	if err != nil {
		return err
	}
	encrypted, err := PassManager.Encrypt(password, b)
	if err != nil {
		return err
	}
	// do backup of current vault
	cleanup, err := PassManager.TmpBackup()
	if err != nil { // backup failed to be created, abort writing
		return err
	}
	// write new vault to file afer backup is done
	if err := PassManager.Write(encrypted); err != nil {
		return err
	}
	// delete tmp backup after nedw vault is written to FS
	if err := cleanup(); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolVar(&override, "override", false, "this will override the password currently stored under this provided account")
}
