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

		var vault map[string]string
		oldPassword, err := clIO.Password()
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
		vault, err = PassManager.Decrypt(oldPassword, fileContent)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		password, err := clIO.SetNewPassword(PassManager.EvaluatePassword)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
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
	rootCmd.AddCommand(passwordCmd)
}
