FROM golang:1.19 AS build 
WORKDIR /app 
COPY ./bin/main.go .
RUN go mod init mariadb-backup && go mod tidy && go mod vendor
RUN go build -o main main.go
CMD ["./main"]