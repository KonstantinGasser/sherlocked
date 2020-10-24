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
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/KonstantinGasser/sherlocked/internal"
	"github.com/spf13/cobra"
)

var username string

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "delete a key value pair from the vault",
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

		if _, ok := vault[username]; !ok {
			fmt.Printf("Sorry mate I could not find a match for '%s', run lock list to see if you misstyped\n", username)
			return
		}
		fmt.Printf("Are you sure you want to delete the user acoont '%s'? [Y/n]: ", username)
		reader := bufio.NewReader(os.Stdin)
		yes, _ := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if strings.TrimSpace(yes) == "Y" {
			delete(vault, username)
			vaultslcie, err := json.Marshal(vault)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			if err := internal.EncryptVault(vaultPath, password, vaultslcie); err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("🗑 user account deleted!")
			return
		}

		fmt.Println("user account NOT deleted!")
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
	delCmd.Flags().StringVarP(&username, "user", "u", "", "account name which you want to delete")

}
