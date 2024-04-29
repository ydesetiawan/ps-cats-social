# ps-cats-social
Project Sprint for Cats Social


### What is this repository for? ###

* Quick summary
* Version
* [Learn Markdown](https://bitbucket.org/tutorials/markdowndemo)

### How do I get set up? ###

* Prerequisites
- go version 1.21 or later
- Mysql 5.7 or later
- Docker 20 or later

* Summary of set up
  Setup web server (non-docker)
- cd cmd/production-api
- go run main.go http

* Golang Migrate
- https://www.freecodecamp.org/news/database-migration-golang-migrate/
```
- migrate create -ext sql -dir db/migrations/ -seq create_user
```

```
migrate -path db/migrations/ -database "postgresql://postgres:@localhost:5432/cats_social?sslmode=disable" -verbose up
```

```
migrate -path db/migrations/ -database "postgresql://postgres:@localhost:5432/cats_social?sslmode=disable" -verbose down
```