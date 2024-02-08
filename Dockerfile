FROM golang:1.21-alpine AS build_base

RUN apk add --no-cache git build-base

ENV GO111MODULE=on
ENV GOPRIVATE=github.com/engagerocketco/*

WORKDIR /app

# To import go-sdk
ARG GITHUB_ACCESS_TOKEN
RUN git config --global url."https://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@github.com".insteadOf "https://github.com"

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Enabling dynamic linking, for the dynatrace monitoring
RUN go build -ldflags="-linkmode=external" -o ./out/go-app ./cmd/app/main.go

# Start fresh from a smaller image
FROM alpine:3.9
RUN apk add ca-certificates build-base

ARG USER=deployment
ARG GROUP=deployment

# This is to create a non-privilege user
RUN addgroup -S $GROUP && adduser -S $USER -G $GROUP

RUN mkdir -p /app/docs

COPY --from=build_base /app/out/go-app /home/$USER/app/go-app
COPY --from=build_base /app/docs/swagger.json /home/$USER/app/docs/swagger.json
COPY --from=build_base /app/Makefile /home/$USER/app/Makefile

# only provide read and execute access to directory for non-privilege user
RUN chmod 550 /home/$USER

# set current user as the non-privilege user
USER $USER

# Set working directory for app
WORKDIR /home/$USER/app

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
# and start the server by adding ` serve` arg
CMD ["./go-app"]
