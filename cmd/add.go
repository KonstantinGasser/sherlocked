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
	"os"

	"github.com/KonstantinGasser/sherlocked/internal"
	"github.com/spf13/cobra"
)

var override bool

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new username and password to the key vault",
	Long:  `with the add command you can add a new password mapped to a username or account name`,
	Run: func(cmd *cobra.Command, args []string) {

		isInit, err := internal.CheckVaultInit(vaultPath)
		if err != nil || !isInit {
			fmt.Println(err.Error())
			return
		}

		password, err := internal.InputPassword()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		uname, pass, err := internal.InputCredentials()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		vault, err := internal.DecryptVault(vaultPath, password)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if key, ok := vault[uname]; ok && !override {
			fmt.Printf("🤔 User %s already stroed use add --override (this option is inreversable) or del -user\n", key)
			return
		}
		vault[uname] = pass

		vaultslcie, err := json.Marshal(vault)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if err := internal.EncryptVault(vaultPath, password, vaultslcie); err != nil {
			fmt.Println(err.Error())
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolVar(&override, "override", false, "this will override the password currently stored under this provided account")
}
