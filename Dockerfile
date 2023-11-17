# use   official Golang image
FROM golang:1.21-alpine3.17

WORKDIR /app

COPY . .

RUN go clean --modcache
RUN GOOS=linux go build -o main cmd/main.go

EXPOSE 3333

CMD ["/app/main"]

