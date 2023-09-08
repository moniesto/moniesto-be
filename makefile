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
	make dropdb || true
	make dropdb-test || true
	make createdb
	make createdb-test
	make migrateup
	make migrateup-test

build:
	go build cmd/main.go

run:
	go run cmd/main.go

nodeman:
	nodemon --exec go run cmd/main.go

test:
	go test ./...

test-c:
	go test ./... -cover

test-v:
	go test ./... -v

test-c-out:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

docker-build:
	docker build -t moniesto:latest .

# compose
compose:
	docker compose down
	docker rmi moniesto-be-api || true
	chmod +x wait-for.sh
	chmod +x start.sh
	docker compose up

# run in development mode
docker-run:
	docker run --name moniesto --network moniesto-network -p 8080:8080 -e DB_SOURCE="postgres://root:secret@moniesto-postgres14:5432/moniesto?sslmode=disable" moniesto:latest

# run in release mode
docker-run-release:
	docker run --name moniesto -p 8080:8080 -e GIN_MODE=release moniesto:latest

# create moniesto docker network
create-docker-network:
	docker network create moniesto-network

# db: connect to docker network
connect-network-db:
	docker network connect moniesto-network moniesto-postgres14

# swagger init
swagger:
	chmod +x swagger.sh
	./swagger.sh