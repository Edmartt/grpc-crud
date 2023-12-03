FROM golang:1.19.9-alpine3.18 AS builder

RUN apk add build-base

WORKDIR /grpc-crud

ENV go env -w GO111MODULE=on
ENV CGP_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

ENV PORT=8080

COPY . .

RUN go build ./... && go build

FROM alpine:3.18

RUN apk update upgrade

WORKDIR /grpc-crud

COPY --from=builder /grpc-crud .
RUN chmod +x grpc-test

EXPOSE ${PORT}

CMD ["./grpc-test"]
