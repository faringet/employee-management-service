USER_SSH_KEY := ~/.ssh/id_rsa
DB_MIGRATIONS_URL := git@github.com:engagerocketco/db-migrations.git

install:
	go mod download

run:
	./go-app

build:
	go build -o ./go-app cmd/app/main.go

lint:
	golangci-lint run --timeout 3m

test:
	go test -coverpkg=./... -coverprofile=coverage.out -covermode=atomic -race ./...

test-report: test
	go tool cover -html=coverage.out

test-docker:
	USER_SSH_KEY=$(USER_SSH_KEY) \
	docker compose -f docker-compose.yml -f docker-compose.tests.yml run --rm --build tests

migrate:
	git clone $(DB_MIGRATIONS_URL) /tmp/db-migrations && \
	goose -dir /tmp/db-migrations/appuser up

generate-swagger:
	swag fmt && swag init -g cmd/app/main.go

# sample: make mock dir=./internal/app name=IOrganizationUserRepository
mock:
	cd $(dir) && mockery --name="$(name)" && cd -
