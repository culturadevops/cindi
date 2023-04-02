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
	"strconv"

	"github.com/culturadevops/cindi/models"
	"github.com/spf13/cobra"
)

// updateCmd represents the del command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "modifica secreto por identificador",
	Long: `modifica un elemento dado su id solo modificara los que son de tipo credenciales
	ex> update -i 17 nuevacontrasenia
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if flags, _ := cmd.Flags().GetBool("id"); flags {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			models.VarSecret.UpdateEncry(Owner, id, args[1])
		}

	},
}

func init() {
	updateCmd.Flags().BoolP("id", "i", false, "Eliminar por id")

	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
