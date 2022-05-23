FROM golang:1.18-alpine

WORKDIR /

COPY go.mod /
RUN go mod download

COPY . .

CMD ["go", "run", "./cmd/cli/main.go"]