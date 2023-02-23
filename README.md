# Run Types

<details><summary>RUN [you need local development]</summary>

</br>

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

</details>

</br>

<details><summary>RUN [all on docker, just using endpoints]</summary>
</br>

Make sure Docker is installed on your machine and `Docker Daemon` is up. (simply run docker).

## Run with using make

```bash
make compose
```

`OR if it failed to run the make command, you can run them manually by:`

```bash
docker compose down
docker rmi moniesto-be-api || true
chmod +x wait-for.sh
chmod +x start.sh
docker compose up
```

</details>

</br>

## Get Environment Variables

1 - Create file with name `app.env`

2 - Go to [moniesto test environment variable](https://docs.google.com/document/d/1jgmkveKCvKAi9UTUsUfRwLrHdB65s2XM5ofS3iQVCcM/edit?usp=sharing)

3 - Copy the content inside `app.env` file

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

:heavy_check_mark: POST /account/password/send_email -> send password reset email

:heavy_check_mark: POST /account/password/change_password -> verify token & change password

:heavy_check_mark: [need test] POST /account/password/verify_token -> verify token

:heavy_check_mark: PUT /account/password -> change password

:heavy_check_mark: [need test] POST /account/email/send_verification_email -> send email verification email

:heavy_check_mark: [need test] POST /account/email/verify_email -> verify email

:heavy_check_mark: [need test] GET /content/posts?subscribed=true&active=true&limit=10&offset=0

:heavy_check_mark: [need test] GET /content/moniests?subscribed=true&limit=10&offset=0

:heavy_check_mark: [need test] GET /users/:username

- [ ] GET /users/:username/posts?limit=<limit>&offset=<offset>

- [ ] GET /users/:username/posts/:post_id

:heavy_check_mark: [need test] GET /users/:username/subscriptions?limit=<limit>&offset=<offset>

- [ ] GET /users/:username/subscribers?limit=<limit>&offset=<offset>

:heavy_check_mark: [need test + payment update] POST /moniests -> be moniest

:heavy_check_mark: [need test] POST /moniests/posts -> create post

:heavy_check_mark: [need test] PATCH /users/profile -> update user profile info

:heavy_check_mark: [need test] PATCH /account/username -> update username

:heavy_check_mark: [need test + payment update] PATCH /moniests/profile -> update moniest profile + subscription info

:heavy_check_mark: [need test + payment update] POST /moniests/:username/subscribe -> subscribe to a moniest

:heavy_check_mark: [need test + payment update] POST /moniests/:username/unsubscribe -> unsubscribe from a moniest

:heavy_check_mark: [need test] GET /crypto/currencies?name=BTC

:heavy_check_mark: GET /assets/error-codes -> get error codes

:heavy_check_mark: GET /assets/validations -> get validation configs

:heavy_check_mark: GET /assets/configs -> get all configurations

:heavy_check_mark: [need test] POST /feedback -> give feedback

- [ ] Search Moniest -> only username

- [ ] POST /account/card -> Add new card

- [ ] DELETE /account/card/:id -> Delete specific card

- [ ] ...
