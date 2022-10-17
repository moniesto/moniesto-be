# RUN

## 1 - [only once] Run Postgres Container (on Docker)

Make sure `Docker Daemon` is up. (simply run docker).

```bash
make postgres
```

## 2 - [only once] Create DB (do only once)

```bash
make createdb
```

## 3 - [only once/or when needed] Run Migrations

```bash
make migrateup
```

## 4 - [when needed] Generate Go code from Queries

- win:

```bash
docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate
```

- bash

```bash
docker run --rm -v "$(pwd):/src" -w /src kjconroy/sqlc generate
```

## 5 - Run the project

run in live reload mode: (need to install nodemon: `npm install -g nodemon`)

```bash
make run-live
```

OR

run (without live reload):

```bash
make run
```

# Setup / Downloads

### 1 - Download Go (v1.19 latest)

You can donwload from [here.](https://go.dev/dl)

### 2 - Download Docker

You can donwload from [here.](https://www.docker.com)

### 3 - Download SQLC

Go tool for DB, [docs](https://docs.sqlc.dev/en/stable/)

You can donwload from [here.](https://docs.sqlc.dev/en/latest/overview/install.html)

### 4 - DBeaver (Optional)

DB tool

You can donwload from [here.](https://dbeaver.io/download)

### 5 - TablePlus (Optional)

DB tool

You can donwload from [here.](https://tableplus.com)

### 6 - Golang Migrate

Migration tool

You can donwload from [here.](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

---

# Back-end Tech Stack

- Language: Golang

- DB: PostgreSQL

- Go tool for DB: [sqlc](https://docs.sqlc.dev/en/stable/)

- Containerization: Docker

- Cloudinary (image service)
