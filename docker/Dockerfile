FROM golang:1.22 AS builder
WORKDIR /app 

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o ./app ./cmd

FROM ubuntu 
WORKDIR /var/www/chrisweir
COPY ./assets ./assets
COPY --from=builder /app/app ./server
CMD ["./server"]

