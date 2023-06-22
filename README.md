# go-api-templete
api template used by go

# How to use

## Setting
1. Clone this repository
2. Execute ```docker compose up -d```
3. Execute ```docker compose exec app go mod init api```
4. Execute ```docker compose exec app go get github.com/gin-gonic/gin```

## Run
1. Edit app/main.go file
2. Execute ```docker compose exec app go run main.go```
3. Access http://localhost:3150/
