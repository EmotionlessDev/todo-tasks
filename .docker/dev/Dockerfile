FROM golang:1.23.1

WORKDIR /api

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /api/bin/main ./cmd/api/

RUN chmod +x /api/bin/main

CMD ["/api/bin/main"]

