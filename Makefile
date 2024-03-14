DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

network: 
		docker network create bank-network

postgres: 
		docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb: 
		docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb: 
		docker exec -it postgres dropdb simple_bank

migrateup:
		migrate -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -path db/migration -verbose up

migratedown:
		migrate -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -path db/migration -verbose down

sqlc:
	 sqlc generate
	 
.PHONY: network postgres createdb dropdb migrateup migratedown start