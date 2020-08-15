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
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/culturadevops/cindi/cmd"
	homedir "github.com/mitchellh/go-homedir"
)

func createfile() string {

	return `
[default]
host = ""
port = "3306"
database = ""
user = ""
password =""
charset = "utf8"
sql_log = false
owner="root"

MaxIdleConns = 10 #空闲时最大的连接数
MaxOpenConns = 100 #最大的连接数 `
}
func crearDirectorioSiNoExiste(directorio string) {
	if _, err := os.Stat(directorio); os.IsNotExist(err) {
		err = os.MkdirAll(directorio, 0755)

		if err != nil {
			// Aquí puedes manejar mejor el error, es un ejemplo
			panic(err)
		}
	}
}
func CrearArchivo(rutaDestino string, data string) {
	err := ioutil.WriteFile(rutaDestino, []byte(data), 0755)
	if err != nil {
		panic(err)
	}
}

func main() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
	dir := home + "/.config/cindi"
	file := dir + "/mysql1.toml"
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Println("Creando archivo de configuración")
		crearDirectorioSiNoExiste(dir)
		CrearArchivo(file, createfile())
		fmt.Println("Necesita establecer la configuración de la base de datos en el archivo:")
		fmt.Println(file)
		os.Exit(1)
	}

	cmd.Execute()
}
