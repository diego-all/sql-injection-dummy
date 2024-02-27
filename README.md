# Dummy API  
    
API written in Golang with SQL injection vulnerability and level code mitigation.

## Run Database

    docker-compose up -d
    docker-compose down

    CURRENT_UID=$(id -u):$(id -g) docker-compose up (colima in Mac OS) 

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
  -e POSTGRES_DB=dbname \
  -p 54325:5432 \
  postgres:12.5-alpine
```
    docker exec -it security-dbs -U user dbname
    docker exec -it sqli-eafit psql -U postgres -d sqli


## SQL injection scenario


Error based sql Injection

    https://go.dev/doc/database/sql-injection

    https://go.dev/doc/database/querying

    https://mariocarrion.com/2021/10/22/golang-software-architecture-security-databases-sql-injection-permissions.html


### Payloads

Non Compliant code

    localhost:9090/vulnerable/users?id=17' OR ''='

Compliant code

    localhost:9090/users?id=46' OR ''='


## Docs

There are environment and collection for Postman in docs folder to facilitate play with API.



















