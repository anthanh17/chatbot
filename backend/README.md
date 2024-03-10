# Backend chatbot

## Require

1. Golang
2. Docker
3. Make
4. Postgres database
5. Pgvector

## How to run

> Update value app.env

Run container database

```
make postgres
```

Create database

```
make createdb
```

Run server

```
make server
```

## How to test

> Check file: `testAPI.http` in directory backend.
