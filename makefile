PORT=5432

postgres:
	docker run --name moniesto-postgres14 -p $(PORT):5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it moniesto-postgres14 createdb --username=root --owner=root moniesto

createdb-test:
	docker exec -it moniesto-postgres14 createdb --username=root --owner=root moniesto-test

dropdb:
	docker exec -it moniesto-postgres14 dropdb moniesto

dropdb-test:
	docker exec -it moniesto-postgres14 dropdb moniesto-test

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:$(PORT)/moniesto?sslmode=disable" -verbose up

migrateup-test:
	migrate -path db/migration -database "postgresql://root:secret@localhost:$(PORT)/moniesto-test?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:$(PORT)/moniesto?sslmode=disable" -verbose down

migratedown-test:
	migrate -path db/migration -database "postgresql://root:secret@localhost:$(PORT)/moniesto-test?sslmode=disable" -verbose down

sqlc:
	sqlc generate

sqlc-bash:
	docker run --rm -v "$(pwd):/src" -w /src kjconroy/sqlc generate

sqlc-win:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

# usage example: make create_migration name=init_schema
create_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

resetdb:
	make dropdb
	make dropdb-test
	make createdb
	make createdb-test
	make migrateup
	make migrateup-test

run:
	go run cmd/main.go

run-live:
	nodemon --exec go run cmd/main.go --signal SIGTERM

test:
	go test ./...

test-c:
	go test ./... -cover

test-v:
	go test ./... -v

test-c-out:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

.PHONY: postgres createdb dropdb migrateup migratedown sqlc sqlc-bash sqlc-win create_migration resetdb run