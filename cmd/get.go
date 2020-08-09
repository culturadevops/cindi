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
	"github.com/culturadevops/cindi/models"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Obtienes el user/contraseña en el portapapeles; puedes usar ctrl+v para pegarla",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		var secret = models.Secret{}
		secret,_=	secret.Get(args[0])
		if flags,_:=cmd.Flags().GetBool("cuenta");flags {
			clipboard.WriteAll(secret.Account);
			fmt.Printf("use ctrl+v para pegar el USUARIO=%v de %v \n" ,secret.Account,args[0])
		}else{
			clipboard.WriteAll(secret.Password);
			fmt.Println("use ctrl+v para pegar la CONTRASEÑA de " +args[0])
		}

		
		
	},
}

func init() {
	getCmd.Flags().BoolP("cuenta", "c", false, "Obtiene el nombre de la cuenta")
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
