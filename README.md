## Requirements

- Go v1.17+
- Docker(include docker-compose)
- Auth0 Account

## Development

### config

```shell
cp backend/config/development/config.toml backend/
```

Set Auth0 properties

### DB Setup

Up to MySQL container

```shell
docker-compose up -d
```

DB migration

```shell
cd backend
go run cmd/migration/main.go
```

Create seed data

```shell
cd backend
go run cmd/seed/main.go
```

### Run

```shell
cd backend
go run main.go
```

On the other window

```shell
cd frontend
yarn dev
```

and open `http://localhost:3000` in your browser, login according to the page
