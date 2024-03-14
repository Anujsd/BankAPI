You need to install

- go
- docker
- makefile
  `sudo apt-get install build-essential`

# DB Handling

### 1. Create Tables using dbdiagram.io

https://dbdiagram.io/d/BankSchema-65eafda8b1f3d4062c73a88b

### 2. Create Postgres docker image

using postgres docker container as db.
you can explore that db using pgadmin or tablePlus

### 3. Create makefile to automate commands

### 4. Use migrate for db migrations

https://github.com/golang-migrate/migrate

Installation steps
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#linux-deb-package

#### Migration Commands

```
migrate create -ext sql -dir db/migrations -seq init_schema

migrate -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -path db/migrations -verbose up

migrate -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -path db/migrations -verbose down
```

### 5. Use sqlc to generate type-safe code from SQL

https://github.com/sqlc-dev/sqlc

Installation https://docs.sqlc.dev/en/latest/overview/install.html

```
sqlc generate
```

### 6. Write test cases
