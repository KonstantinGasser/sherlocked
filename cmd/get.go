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

	"github.com/KonstantinGasser/sherlocked/internal"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var name string
var hidePassword bool
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get returns the password stored for a given account",
	Long:  ``,
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

		vault, err := internal.DecryptVault(vaultPath, password)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if key, ok := vault[name]; !ok {
			fmt.Printf("😐 No user %s in the vault", key)
			return
		}
		clipboard.WriteAll(vault[name])
		if !hidePassword {
			fmt.Print(vault[name])
		}

		return
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&name, "user", "u", "", "Specifies the account under which the password is stored")
	getCmd.Flags().BoolVar(&hidePassword, "hide", false, "only write password to clipboard")
}
