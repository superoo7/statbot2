FROM golang:1.11-alpine AS build_base
RUN apk add bash ca-certificates git gcc g++ libc-dev
ADD . /go/src/app
WORKDIR /go/src/app
ENV GO111MODULE=on
RUN go mod download
RUN go build
RUN ./statbot2