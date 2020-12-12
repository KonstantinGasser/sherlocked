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

		if err := runDel(args); err != nil {
			fmt.Println(err.Error())
			return
		}
		//
		//
	},
}

// run func holds the logic for the password command
// is a separated function to test the code proper
func runDel(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("😐 No user specified")
	}

	username := args[0]
	password, err := clIO.Password()
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

	if _, ok := vault[username]; !ok {
		return fmt.Errorf("Sorry mate I could not find a match for '%s'\nrun lock list to see if you misstyped the account name\n", username)
	}

	fmt.Printf("Are you sure you want to delete the user acoont '%s'? [Y/n]: ", username)
	reader := bufio.NewReader(os.Stdin)
	yes, _ := reader.ReadString('\n')
	if err != nil {
		return err
	}
	if strings.TrimSpace(yes) != "Y" {
		return fmt.Errorf("user account NOT deleted!")

	}

	// delete user account from vault
	delete(vault, username)
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
	fmt.Println("🗑 user account deleted!")
	return nil
}

func init() {
	rootCmd.AddCommand(delCmd)
}
