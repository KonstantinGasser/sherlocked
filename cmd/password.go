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

// passwordCmd represents the password command
var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "set the password to encrypt and decrypt your vault",
	Long:  `Choose a strong password! And dont forget it - If you choose to go that way..I can't follow you nor can I help you`,
	Run: func(cmd *cobra.Command, args []string) {

		// check if vault exists
		// yes: decrypt vault encrypt with new password
		// no: create new vault encrypt with new password
		fmt.Println(isInit)
		var vault map[string]string
		if isInit { // vault exists
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
		}
		if !isInit { // create new vault
			vault = make(map[string]string)
		}

		password1, err := clIO.SimpleText("New Password: ")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		passwordStrength := PassManager.EvaluatePassword(password1)
		if passwordStrength < 50 {
			fmt.Print("Mhm looks like this is not the best password..😅 - try again [Y/n]: ")
			reader := bufio.NewReader(os.Stdin)
			tryAgain, _ := reader.ReadString('\n')
			if strings.TrimSpace(tryAgain) == "Y" {
				password1, err = clIO.SimpleText("😏 choose wisely: ")
			}
			fmt.Print("\n")
		}

		fmt.Println("🙃 Just to make sure...confirm your password")
		password2, err := clIO.SimpleText("Repeat Password: ")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if password1 != password2 {
			fmt.Println("They don't match let's do it again, shall we? 🤦🏼‍♀️")
			return
		}
		// write changed vault
		b, err := PassManager.Serialize(vault)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		encrypted, err := PassManager.Encrypt(password1, b)
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

		// var vault map[string]string
		// var password1 string
		// var password2 string
		//
		// if isInit {
		// 	password, err := internal.InputText("Current Password: ")
		// 	if err != nil {
		// 		fmt.Println(err.Error())
		// 		return
		// 	}
		// 	vault, err = internal.DecryptVault(vaultPath, password)
		// 	if err != nil {
		// 		fmt.Println(err.Error())
		// 		return
		// 	}
		// }
		//
		// if !isInit {
		// 	vault = make(map[string]string)
		// }
		//
		// password1, err := internal.InputText("Password: ")
		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	return
		// }
		// passwordStrength := internal.EvaluatePassword(password1)
		// if passwordStrength < 50 {
		// 	fmt.Print("Mhm looks like this is not the best password..😅 - try again [Y/n]: ")
		// 	reader := bufio.NewReader(os.Stdin)
		// 	tryAgain, _ := reader.ReadString('\n')
		// 	if strings.TrimSpace(tryAgain) == "Y" {
		// 		password1, err = internal.InputText("😏 choose wisely: ")
		// 	}
		// 	fmt.Print("\n")
		// }
		//
		// fmt.Println("🙃 Just to make sure...confirm your password")
		// password2, err = internal.InputText("Password: ")
		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	return
		// }
		// if password1 != password2 {
		// 	fmt.Println("They don't match let's do it again, shall we? 🤦🏼‍♀️")
		// 	return
		// }
		//
		// vaultslcie, err := json.Marshal(vault)
		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	return
		// }
		// if err := internal.EncryptVault(vaultPath, password1, vaultslcie); err != nil {
		// 	fmt.Println(err.Error())
		// 	return
		// }
		// fmt.Println("Vault is now encrypted with the new password ✅")
	},
}

func init() {
	rootCmd.AddCommand(passwordCmd)

}
