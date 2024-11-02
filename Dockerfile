FROM golang:1.20-alpine AS builder

WORKDIR C:\Users\Predator\go\src

RUN apk update && apk --no-cache add bash gcc gettext

COPY ./go.mod ./go.sum ./

RUN go mod download