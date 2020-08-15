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
	Short: "Agrega una secreto, los modificadores definen el tipo",
	Long: `Al usar este comando podras guardar un secreto de 
	tipo token un solo parametro
	tipo credenciales dos parametros
	tipo amazon tres parametros
Recuerda el ultimo parametro siempre sera el secreto
			`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if flags, _ := cmd.Flags().GetBool("secret1"); flags {
			if len(args) == 2 {
				return models.VarSecret.Additem(libs.Owner, args[0], args[1])
			} else {
				print("falta un valor 'secret'")
				return nil
			}
		}
		if flags, _ := cmd.Flags().GetBool("secret2"); flags {
			if len(args) == 3 {
				return models.VarSecret.Additem1(libs.Owner, args[0], args[1], args[2])
			} else {
				print("Necesitas 'idenficador' 'user' 'secret' ")
				return nil
			}

			//	return models.VarCredential.Add(libs.Owner, args[0], args[1], args[2])
		} else if flags, _ := cmd.Flags().GetBool("secret3"); flags {
			if len(args) == 4 {
				return models.VarSecret.Additem2(libs.Owner, args[0], args[1], args[2], args[3])
			} else {
				print("Necesitas 'idenficador' 'account' 'user' 'secret' ")
				return nil
			}
		}
		print("Falta modificador user -h para ver opciones")
		return nil
	},
}

func init() {
	addCmd.Flags().BoolP("secret1", "t", false, "Crear un secreto en forma token de 'secret' ")
	addCmd.Flags().BoolP("secret2", "c", false, "Crear un secreto en forma credencial de user/secret")
	addCmd.Flags().BoolP("secret3", "a", false, "Crear un secreto en forma amazon credencial de Accountid/user/secret")
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
