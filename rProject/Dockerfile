FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o message-broker

EXPOSE 8080

CMD ./message-broker