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
	"strconv"

	"github.com/atotto/clipboard"
	"github.com/culturadevops/cindi/models"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Obtienes el secreto en el portapapeles; puedes usar ctrl+v para pegarla",
	Long: `al usar este comando podras obtener el secreto del registro 
	siempre obtendras el ultimo parametro en el portapapeles
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var x models.Items
		secret := models.Secret{}

		if flags, _ := cmd.Flags().GetBool("id"); flags {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			secret, _ = models.VarSecret.GetForId(Owner, id)
		} else {
			secret, _ = models.VarSecret.Get(Owner, args[0])

		}
		json.Unmarshal([]byte(secret.Secret), &x)
		if x.Type == "credential" {
			fmt.Printf("Identificador %v - %v \nUser= %v \n", args[0], secret.Name, x.Items["user"])
		}
		if x.Type == "amazon" {
			fmt.Printf("Identificador %v - %v \nAccount ID= %vUser= %v \n", args[0], secret.Name, x.Items["account"], x.Items["user"])
		}
		if flags, _ := cmd.Flags().GetBool("out"); flags {
			fmt.Printf("%v", x.Items["secret"])
		} else {
			clipboard.WriteAll(x.Items["secret"])
			fmt.Printf("use ctrl+v para pegar el SECRETO de %v - %v  \n", args[0], secret.Name)
		}

	},
}

func init() {
	getCmd.Flags().BoolP("id", "i", false, "Obtien una secret dado el numero de id")
	getCmd.Flags().BoolP("out", "o", false, "Obtien un secreto por pantalla ideal para ssh y command ")
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
