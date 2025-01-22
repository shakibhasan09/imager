FROM golang:1.23.0 AS build

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/server/main.go

FROM debian:bookworm-slim

WORKDIR /usr/src/app

COPY --from=build /usr/src/app/main ./

VOLUME [ "/usr/src/app/mount" ]

EXPOSE 8080

CMD ["./main"]
