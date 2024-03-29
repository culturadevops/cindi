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

	"github.com/culturadevops/cindi/models"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lista todos los secretos",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var x models.Items
		listsecret := models.VarSecret.List(Owner)

		fmt.Printf("Hay %v secretos\n", len(listsecret))
		for _, value := range listsecret {
			json.Unmarshal([]byte(value.Secret), &x)
			if x.Type != "command" {
				fmt.Printf("%v-%v tipo:%v\n", value.ID, value.Name, x.Type)

			} else {
				fmt.Printf("%v-%v t:%v %v\n", value.ID, value.Name, x.Type, x.Items["secret"])
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
