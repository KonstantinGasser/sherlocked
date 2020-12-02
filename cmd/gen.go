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
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var (
	length    int
	upperCase int
	lowerCase int = 2
	numbers   int
	specials  int
	ignore    string

	newUser string
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		randPassword := internal.GeneratePassword(length, upperCase, lowerCase, numbers, specials, ignore)

		// if newUser == "nil" {
		// 	fmt.Printf("🤨 Seems like you forgot to tell me which user that password is for?")
		// 	return
		// }

		if newUser != "nil" {
			// fetch vault password and decrypt vault
			password, err := internal.InputPassword()
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			vault, err := internal.DecryptVault(vaultPath, password)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			// add user and password to vault
			if _, ok := vault[newUser]; ok {
				fmt.Printf("🤔 User %s already stroed use add --override (this option is inreversable) or del -user\n", newUser)
				return
			}
			vault[newUser] = randPassword

			// encrypt and write vault
			vaultslcie, err := json.Marshal(vault)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			if err := internal.EncryptVault(vaultPath, password, vaultslcie); err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("👽 User: %s with generate password added to vault!\n", newUser)
		}

		clipboard.WriteAll(randPassword)
		fmt.Printf("🗝  %s\n", randPassword)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.Flags().StringVarP(&newUser, "create", "C", "nil", "when given password and user will be added to the vault")
	genCmd.Flags().IntVarP(&length, "length", "l", 8, "determin the length of the generated password (default=8)")
	genCmd.Flags().IntVarP(&upperCase, "uppers", "u", 2, "determin the number of upper case chars")
	genCmd.Flags().IntVarP(&numbers, "numbers", "n", 2, "determin the number of numbers")
	genCmd.Flags().IntVarP(&specials, "specials", "s", 2, "determin the number of special chars (+_-?.@#$%!)")
	genCmd.Flags().StringVar(&ignore, "ignore", "i", "ignore characters (char_1,char_2,char_n...)")

}
