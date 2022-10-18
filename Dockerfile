FROM golang:1.19.1-alpine

WORKDIR /app

COPY . .

RUN go build -o /build

CMD ["/build"]