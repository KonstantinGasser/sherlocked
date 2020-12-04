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
	"strings"

	"github.com/KonstantinGasser/sherlocked/internal"
	"github.com/spf13/cobra"
)

var PassManager internal.PasswordManager
var clIO internal.IO
var version = "v1.0.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lock",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initVault() {
	homeDir, _ := os.UserHomeDir()
	vaultPath := strings.Join([]string{homeDir, ".sherlocked"}, "/")

	clIO = internal.NewIO()
	PassManager = internal.NewPasswordManager(vaultPath)
	if ok, _ := PassManager.IsInit(); !ok {
		if err := PassManager.Init(clIO); err != nil {
			fmt.Println(err.Error())
		}
		os.Exit(1)
	}
}

func init() {}
