# syntax=docker/dockerfile:1

# Build stage

FROM golang:1.18.3-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /identity-check

# Run Stage

FROM alpine:3.16

WORKDIR /app
COPY --from=builder  identity-check .

EXPOSE 9000

CMD [ "./identity-check" ]