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
	"github.com/culturadevops/cindi/libs"
	"github.com/culturadevops/cindi/models"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Agrega una creadencial",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if flags, _ := cmd.Flags().GetBool("secret"); flags {
			return models.VarSecret.Add(libs.Owner, args[0], args[1])
		} else {
			return models.VarCredential.Add(libs.Owner, args[0], args[1], args[2])
		}

	},
}

func init() {
	addCmd.Flags().BoolP("secret", "s", false, "Crear un secreto en forma de token")
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
