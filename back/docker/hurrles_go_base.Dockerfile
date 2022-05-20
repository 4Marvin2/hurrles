FROM golang:1.17.2

WORKDIR /app

COPY . .

RUN make build-go