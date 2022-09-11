FROM golang:1.19-alpine3.16

WORKDIR /app

ENV GO111MODULE=on

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .

RUN go build -o bin/server cmd/server/main.go

CMD ["./bin/server"]