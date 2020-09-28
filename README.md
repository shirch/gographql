# gographql

## Steps to run the application:

1. docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=dbpass -e MYSQL_DATABASE=dbname -d mysql:latest
2. go run server.go
