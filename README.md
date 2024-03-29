## About The Project
Cindi C.L.I una aplicación que permite guardar credenciales de tipo usuario/pass de forma rapida y te permite obtenerlas en el portapapeles del computador dado un identificador unico

La contraseña nunca sera mostrada en pantalla pasara de la aplicacion a la memoria y podras usarla presionando Ctrl+v asi podras pegarla directamente en los inputs secret/password de las aplicaciones donde las necesites

Nota: funciona en windows y linux

## Usage

### Agregar credencial 
para agregar solo debes indicar los siguiente
```
cindi add identificadorunico  micorreo@domini.com  misuperpassword 
```

### Obtener Credencial 
para obtener solo debes indicar los siguiente
```
cindi get identificadorunico  
```

### Otros 
Tambien puedes listar eliminar y modificar las credenciales y para ver todo los comandos puedes usar

```
cindi help
```
para ver la ayuda de los comandos internos puede usar el comando seguido del modificador -h 

```
cindi add -h  
```

## Pre Requisito para usar

### Windows 
1. Cindi.exe en alguna carpeta de tu preferencia
2. Configurar las variables de entorno la ruta del cindi.exe (opcional)
3. Mysql server

### Linux 
1. Cindi en alguna carpeta de tu preferencia
2. Configurar las variables de entorno la ruta del cindi
3. Mysql server

## mysql server
docker run -d -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=secret mysql:5.7.40-debian
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' 59259d6b98da0bd27116d252bcd1ea1aa501150c3ed46e3b0724d73728ca2c9a
docker exec -it mysql-db mysql -p

## Installation

1. Ejecuta el query de sql que puede encontrar en copi/sql/tablas.sql en tu msyql server
2. Copiar el archivo copi/cindi/mysql.toml  en C:\Users\tu-usuario\.config\cindi\mysql.toml
3. Cambiar la configuracion de mysql.toml a la configuracion de tu mysql server

## Compilar cindi en linux

para compilar el proyecto debes descargartelo o hacer git clone del proyecto, entrar en el y escribir go install y con esto ya tendras cindi en la ruta go/bin

nota: si al ejecutar cindi no puedes ver el aplicativo debes configurar la ruta de go/bin en las variables de entorno

```
export PATH=$PATH:/usr/local/go/bin
```
## Compilar cindi en windows
para compilar el proyecto debes descargartelo o hacer git clone del proyecto, entrar en el y escribir go install y con esto ya tendras cindi en la ruta go/bin

nota: si al ejecutar cindi no puedes ver el aplicativo debes configurar la ruta de go/bin en las variables de entorno

```
export PATH=$PATH:/c/usr/local/go/bin
```
nota esto funciona en cualquier consola como git bash / vscode go no funciona en power shell




# Mis Libros:

[![libros futuro es devops ](https://github.com/culturadevops/recursos/blob/master/portada-futuro-es-devops.png)](https://amzn.to/3S8AGG9) [![libros herramientas devops](https://github.com/culturadevops/recursos/blob/master/portada-herramientasdevops.png)](https://amzn.to/3ga1c4E)

# Mi canal de cultura Devops

[![canal de youtube sobre devops ](https://github.com/culturadevops/recursos/blob/master/logo-culturadevops.png)](https://www.youtube.com/channel/UCfJ67eVA7DkKbbIF5ceJDMA?sub_confirmation=1) 
