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
	"strconv"

	"github.com/atotto/clipboard"
	"github.com/culturadevops/cindi/libs"
	"github.com/culturadevops/cindi/models"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Obtienes el user/contraseña en el portapapeles; puedes usar ctrl+v para pegarla",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if flags, _ := cmd.Flags().GetBool("secret"); flags {
			secret := models.Secret{}
			if flags, _ := cmd.Flags().GetBool("id"); flags {
				id, _ := strconv.ParseInt(args[0], 10, 64)
				secret, _ = models.VarSecret.GetForId(libs.Owner, id)
			} else {
				secret, _ = models.VarSecret.Get(libs.Owner, args[0])

			}
			clipboard.WriteAll(secret.Secret)
			fmt.Printf("use ctrl+v para pegar el SECRETO de %v \n", args[0])
		} else {
			credential := models.Credential{}
			if flags, _ := cmd.Flags().GetBool("id"); flags {
				id, _ := strconv.ParseInt(args[0], 10, 64)
				credential, _ = models.VarCredential.GetForId(libs.Owner, id)
			} else {
				credential, _ = models.VarCredential.Get(libs.Owner, args[0])
			}
			if flags, _ := cmd.Flags().GetBool("user"); flags {
				clipboard.WriteAll(credential.Account)
				fmt.Printf("use ctrl+v para pegar el USUARIO=%v de %v \n", credential.Account, credential.Name)
			} else {
				clipboard.WriteAll(credential.Password)
				fmt.Println("use ctrl+v para pegar la CONTRASEÑA de " + credential.Name)
			}
		}

	},
}

func init() {
	getCmd.Flags().BoolP("id", "i", false, "Obtiene una cuenta dado el numero de id")
	getCmd.Flags().BoolP("user", "u", false, "Obtiene el nombre de la cuenta")
	getCmd.Flags().BoolP("secret", "s", false, "Crear un secreto en forma de token")
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
