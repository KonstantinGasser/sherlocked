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
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "delete a key value pair from the vault",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Verify that a password is set for the vault
		// and a default vault exists
		initVault()
		if len(args) < 1 {
			fmt.Println("😐 No user specified")
			return
		}

		username := args[0]
		password, err := clIO.Password()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
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

		if _, ok := vault[username]; !ok {
			fmt.Printf("Sorry mate I could not find a match for '%s'\nrun lock list to see if you misstyped the account name\n", username)
			return
		}

		fmt.Printf("Are you sure you want to delete the user acoont '%s'? [Y/n]: ", username)
		reader := bufio.NewReader(os.Stdin)
		yes, _ := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if strings.TrimSpace(yes) != "Y" {
			fmt.Println("user account NOT deleted!")
			return
		}

		// delete user account from vault
		delete(vault, username)
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
		after, err := PassManager.Backup(func() error {
			return PassManager.Write(encrypted)
		})
		if err != nil { // backup failed to be created, abort writing
			fmt.Println(err.Error())
			return
		}
		if err := after(); err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("🗑 user account deleted!")
		return
		//
		//
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
