# How to install PostgreSQL
https://www.calhoun.io/how-to-install-postgresql-9-5-on-ubuntu-16-04/
# Installation
$ sudo apt-get update
$ sudo apt-get install postgresql postgresql-contrib
$ sudo -u postgres psql
# Setting up a password for the postgres role
$ sudo -u postgres psql
$ ALTER USER postgres WITH ENCRYPTED PASSWORD '1'; 
# Creating a Postgres database
$ psql -U postgres
$ CREATE DATABASE quang_database

===========================================================
# Gen proto file 
make gen-proto

# RESTovergRPC
Create a directory services with rest over grpc

# Testing local

Run postgres
============
```console
$ docker run --rm -d --name postgres -p 1234:1234 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgres postgres
```

Run server
==========

```
$ go run main.go
```
=============
Run test API
# test echo API
curl -X POST "http://localhost:8081/api/v1/echo/test"
# addEntry to SQL
curl -X POST -d'{"directory_name": "test", "entry":{"name": "abc","last_name": "cde","ph_number": "090123456"}}' http://localhost:8081/api/v1/addEntry
# get User API to get the first User's database
curl -X POST -d'{"command": "First"}' http://localhost:8081/api/v1/getUser
=============
