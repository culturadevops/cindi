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
	//"log"
	"os"

	"github.com/culturadevops/GORM/libs"

	//"github.com/culturadevops/cindi/libs"
	"github.com/culturadevops/cindi/models"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	config "github.com/spf13/viper"
)

var Owner string

//var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cindi",
	Short: "Getionador de credenciales v1.1",
	Long:  ``,
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

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cindi.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dbConfig := libs.Configure(home, ".config/cindi/mysql")
	libs.DB = dbConfig.InitMysqlDB()
	Owner = config.GetString("default.owner")
	// Search config in home directory with name ".cindi" (without extension).
	//viper.AddConfigPath(home)
	//viper.SetConfigName(".cindi")

	//viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	//viper.SetConfigName(".config/cindi/mysql")
	//file := home + ".config/cindi/mysql"

	/*	if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}
		Owner = viper.GetString("default.owner")

		if viper.GetString("default.host") == "" {
			fmt.Println("falta host en el archivo " + file)
			os.Exit(1)
		}
		if viper.GetString("default.database") == "" {
			fmt.Println("falta database en el archivo de config " + file)
			os.Exit(1)
		}
		if viper.GetString("default.user") == "" {
			fmt.Println("falta user en el archivo de config " + file)
			os.Exit(1)
		}
		if viper.GetString("default.password") == "" {
			fmt.Println("falta password en el archivo de config " + file)
			os.Exit(1)
		}
		dbConfig := libs.DbConfig{
			viper.GetString("default.host"),
			viper.GetString("default.port"),
			viper.GetString("default.database"),
			viper.GetString("default.user"),
			viper.GetString("default.password"),
			viper.GetString("default.charset"),
			viper.GetInt("default.MaxIdleConns"),
			viper.GetInt("default.MaxOpenConns"),
		}
		if viper.GetBool("default.sql_log") {
			libs.DB.LogMode(true)
		} else {
			libs.DB.LogMode(false)
		}

	*/

	//libs.DB = dbConfig.InitDB()
	//Owner = viper.GetString("default.owner")
	models.VarSecret = &models.Secret{}

}
