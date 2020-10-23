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
	"encoding/json"
	"fmt"

	"github.com/KonstantinGasser/sherlocked/internal"
	"github.com/spf13/cobra"
)

// passwordCmd represents the password command
var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "set the password to encrypt and decrypt your vault",
	Long:  `Choose a strong password! And dont forget it - If you choose to go that way..I can't follow you nor can I help you`,
	Run: func(cmd *cobra.Command, args []string) {
		isInit, _ := internal.CheckVaultInit(vaultPath)

		var vault map[string]string
		var password1 string
		var password2 string

		if isInit {
			password, err := internal.InputNewPassword("Current Password: ")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			vault, err = internal.DecryptVault(vaultPath, password)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}

		if !isInit {
			vault = make(map[string]string)
		}

		password1, err := internal.InputNewPassword("Password: ")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("🙃 Just to make sure...confirm your password")
		password2, err = internal.InputNewPassword("Password: ")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if password1 != password2 {
			fmt.Println("They don't match let's do it again, shall we? 🤦🏼‍♀️")
			return
		}

		vaultslcie, err := json.Marshal(vault)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if err := internal.EncryptVault(vaultPath, password1, vaultslcie); err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Vault is now encrypted with the new password ✅")
	},
}

func init() {
	rootCmd.AddCommand(passwordCmd)

}
