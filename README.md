# ps-cats-social
Project Sprint for Cats Social


### What is this repository for? ###

* Quick summary
* Version
* [Learn Markdown](https://bitbucket.org/tutorials/markdowndemo)

### How do I get set up? ###

#### Prerequisites
- go version 1.21 or later
- PostGres v16

#### Summary of set up and run
- cd cmd/api
- go run main.go http

#### Golang Migrate
- https://www.freecodecamp.org/news/database-migration-golang-migrate/

### Run Migration using MakeFile

- make migration_setup
- make migration_up
- make migration_down
- make migration_fix