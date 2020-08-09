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
	
	"github.com/culturadevops/cindi/models"
	"github.com/spf13/cobra"

	"strconv"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Elimina credencial dando un identificador",
	Long: `Elimina un secreto dando un identificador puede ser el nombre o el id del registro`,
	Run: func(cmd *cobra.Command, args []string) {
		var secret = models.Secret{}
		if flags,_:=cmd.Flags().GetBool("id");flags {
			id,_:=strconv.ParseInt(args[0],10, 64)
			secret.DelForId(id)
		}else{
			secret.Del(args[0])
		}
		
	},
}

func init() {
	delCmd.Flags().BoolP("id", "i", false, "Eliminar por id")
	rootCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
