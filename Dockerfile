FROM golang:alpine as builder
RUN mkdir -p /assets
ADD fly/fly-*-linux-amd64.tgz /assets/

COPY concourse-team-resource /go/src/github.com/dmechas/concourse-team-resource

ENV CGO_ENABLED 0
RUN go build -o /assets/out github.com/dmechas/concourse-team-resource/cmd/out

FROM alpine:edge AS resource
RUN apk add --no-cache bash tzdata ca-certificates
COPY --from=builder assets/ /opt/resource/
RUN chmod +x /opt/resource/*

FROM resource
