monta un servidor y agrega la tablas.sql

## mysql server
docker run -d -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=secret mysql:5.7.40-debian
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' mysql
docker exec -it mysql-db mysql -p


# Mis Libros:

[![libros futuro es devops ](https://github.com/culturadevops/recursos/blob/master/portada-futuro-es-devops.png)](https://amzn.to/3S8AGG9) [![libros herramientas devops](https://github.com/culturadevops/recursos/blob/master/portada-herramientasdevops.png)](https://amzn.to/3ga1c4E)

# Mi canal de cultura Devops

[![canal de youtube sobre devops ](https://github.com/culturadevops/recursos/blob/master/logo-culturadevops.png)](https://www.youtube.com/channel/UCfJ67eVA7DkKbbIF5ceJDMA?sub_confirmation=1) 
