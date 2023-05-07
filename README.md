## Packages
- gin-gonic -> Web Server
- gorm      -> ORM


## How to
1.0 create a `.env` file using `.env.example` as template

2.1 docker-compose:
```sh
docker-compose up -d --build

docker exec -it categories-api zsh

go mod tidy

go run cmd/api/main.go
```

## Tools:
generate tests:
https://www.airtest.dev/

## Troubleshooting
1. verify db connection
```sh
docker exec -it categories-db bash
su postgres
psql "dbname=meetup host=localhost user=postgres password=password123 port=5432"

\l # list databases

exit
exit
exit
```