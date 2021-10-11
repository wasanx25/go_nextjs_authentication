## Requirements

- Go v1.17+
- Docker(include docker-compose)
- Auth0 Account

## Development

### config

```shell
cp config/development/config.toml .
```

Set Auth0 properties

### DB Setup

Up to MySQL container

```shell
docker-compose up -d
```

DB migration

```shell
go run cmd/migration/main.go
```

Create seed data

```shell
go run cmd/seed/main.go
```

### Run

```shell
go run main.go
```

On the other window

```shell
cd frontend
yarn dev
```

and open `http://localhost:3000` in your browser, login according to the page
