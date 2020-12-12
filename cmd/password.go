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

// passwordCmd represents the password command
var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "set the password to encrypt and decrypt your vault",
	Long:  `Choose a strong password! And dont forget it - If you choose to go that way..I can't follow you nor can I help you`,
	Run: func(cmd *cobra.Command, args []string) {

		// Verify that a password is set for the vault
		// and a default vault exists
		initVault()
		if err := runPassword(); err != nil {
			fmt.Println(err.Error())
			return
		}

	},
}

// run func holds the logic for the password command
// is a separated function to test the code proper
func runPassword() error {
	var vault map[string]string
	oldPassword, err := clIO.Password()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// get encryted vault
	fileContent, err := PassManager.Read()
	if err != nil {
		return err
	}
	// decrypt vault
	vault, err = PassManager.Decrypt(oldPassword, fileContent)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	password, err := clIO.SetNewPassword(PassManager.EvaluatePassword)
	if err != nil {
		return err
	}
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
	rootCmd.AddCommand(passwordCmd)
}
