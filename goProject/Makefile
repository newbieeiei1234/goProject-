sqlc:
	sqlc generate

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simpledb

dropdb:
	docker exec -it postgres12 dropdb simpledb

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret1 -d postgres:16-alpine

migrateup:
	migrate -path db/migration -database "postgresql://root:secret1@localhost:5432/simpledb?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret1@localhost:5432/simpledb?sslmode=disable" -verbose down

server:
	go run main.go

mockdb:
	mockgen -package mockdb -destination db/mock/store.go  goProject/db/sqlc Store


.PHONY: sqlc createdb dropdb postgres migrateup server mockdb