FROM golang:1.19.5

WORKDIR /tests
COPY . .
RUN go mod download
RUN go test ./...
