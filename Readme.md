# Using DB with GO

## Postgres and Mysql with Docker

```sh
#run
docker-compose up -d

#stop
docker-compose stop

```

## Client GUI DB

please install DBeaver(my recommendation)
[dbeaver](https://dbeaver.io/download/)

or we can use pgAdmin for GUI DB

then go to url: [localhost:5050](http://localhost:5050) in the browser

remember to create a database name 'godb' before to run

```sh
#run
go run main.go

```
