FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./
COPY docs ./docs

RUN go mod download

COPY cmd ./cmd
COPY config ./config
COPY internal ./internal
COPY .env ./.env

RUN go install github.com/swaggo/swag/cmd/swag@v1.8.12
RUN swag init --parseDependency --parseInternal --parseDepth 2 -g cmd/app/main.go

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/app/main.go

CMD ./app