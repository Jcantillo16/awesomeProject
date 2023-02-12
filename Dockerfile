FROM golang:latest AS builder

RUN apt-get update
RUN apt-get install -y git

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /go/src/awesomeProject

COPY go.mod .

RUN go mod download

COPY . .

EXPOSE 8080

RUN go build main.go

CMD ["go", "run", "main.go"]

# Path: docker-compose.yml









