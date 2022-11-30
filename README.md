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

# ENDPOINTS

:heavy_check_mark: POST /account/login

:heavy_check_mark: POST /account/register

:heavy_check_mark: GET /account/usernames/:username/check

:heavy_check_mark: [need test] PUT /account/password -> [unauthenticated case] [forget password - send email case]

:heavy_check_mark: [need test] PUT /account/password -> [unauthenticated user] [forget password - change password case]

:heavy_check_mark: PUT /account/password -> [authenticated user] [change password case]

- [ ] GET /content/posts?subscribed=<subscribed>&limit=<limit>&offset=<offset>

- [ ] GET /content/moniests?subscribed=<subscribed>&limit=<limit>&offset=<offset>

- [ ] GET /users/:username

- [ ] GET /users/:username/posts?limit=<limit>&offset=<offset>

- [ ] GET /users/:username/posts/:post_id

- [ ] GET /users/:username/subscriptions?limit=<limit>&offset=<offset>

- [ ] GET /users/:username/subscribers?limit=<limit>&offset=<offset>

:heavy_check_mark: [need test + latest check] POST /moniests

- [ ] POST /moniests/posts

- [ ] PATCH /account/profile

- [ ] PATCH /moniest/subscription-info

- [ ] POST /moniest/:username/subscriptions

- [ ] ...
