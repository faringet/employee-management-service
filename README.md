# go-template
Backend go-template

## Config

### API

##### Required

```
SERVER_HOST - API Server Address, default: localhost
SERVER_PORT - API Server Port, default: 8080  
SERVER_LOG_LEVEL - API Server Log Level, default: development  
```

### DB

##### Required

```
POSTGRES_HOST - Database Host
POSTGRES_PORT - Database Port
POSTGRES_USER - Database User
POSTGRES_PASSWORD - Database Password
POSTGRES_DB - Database Name
```

### NATS

##### Required

```
NATS_HOST - NATS Address
NATS_PORT - NATS Port
```

##### Optional

```
TEST_ENV - bool value, represent is server works on test env or not, default: false
```

# SWAGGER
before push
```shell
swag fmt
```

for regenerate swagger file
```shell
go generate ./...
```

swagger endpoint
> {base_url}/swagger/index.html

# Run app locally

create .env file

```shell
go run cmd/app/main.go
```