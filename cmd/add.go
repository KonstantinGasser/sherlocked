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
	"os"

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

		password, err := clIO.Password()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		uname, pass, err := clIO.Credentials()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// get encryted vault
		fileContent, err := PassManager.Read()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// decrypt vault
		vault, err := PassManager.Decrypt(password, fileContent)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// check if user exists in vault with same key
		if _, ok := vault[uname]; ok && !override {
			fmt.Printf("🤔 User %s already stroed use add --override (this option is inreversable) or del -user\n", uname)
			return
		}

		// add new account to vault
		vault[uname] = pass

		// write changed vault
		b, err := PassManager.Serialize(vault)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		encrypted, err := PassManager.Encrypt(password, b)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// do backup of current vault
		cleanup, err := PassManager.TmpBackup()
		if err != nil { // backup failed to be created, abort writing
			fmt.Println(err.Error())
			return
		}
		// write new vault to file afer backup is done
		if err := PassManager.Write(encrypted); err != nil {
			fmt.Println(err.Error())
			return
		}
		// delete tmp backup after nedw vault is written to FS
		if err := cleanup(); err != nil {
			fmt.Println(err.Error())
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolVar(&override, "override", false, "this will override the password currently stored under this provided account")
}
