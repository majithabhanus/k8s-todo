FROM golang:1.24.3-alpine AS builder

WORKDIR /todo

COPY go.mod go.sum ./
RUN go mod download
COPY . .

ENV DB_USER=admin
ENV DB_PASS=admin123
ENV DB_HOST=mysql-container
ENV DB_PORT=3306
ENV DB_NAME=mydb
ENV JWT_SECRET=G9s!zX7$eA2@tKp

EXPOSE 9000

RUN go build -o todo ./cmd/main.go

FROM alpine:latest

COPY --from=builder /todo/todo .

EXPOSE 9000

ENTRYPOINT [ "./todo" ]