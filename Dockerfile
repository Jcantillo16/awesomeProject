FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY . /app

RUN go mod download

COPY *go ./

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]





