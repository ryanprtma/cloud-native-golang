FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o market-place

EXPOSE 8080

CMD ./market-place