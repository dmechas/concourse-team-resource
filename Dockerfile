FROM golang:alpine as builder
RUN mkdir -p /assets
ARG FLY_VERSION=5.7.2

ADD https://github.com/concourse/concourse/releases/download/v${FLY_VERSION}/fly-${FLY_VERSION}-linux-amd64.tgz /tmp/fly.tgz
RUN tar zxf /tmp/fly.tgz -C /assets/

COPY . /go/src/github.com/dmechas/concourse-team-resource

ENV CGO_ENABLED 0
RUN go build -o /assets/in github.com/dmechas/concourse-team-resource/cmd/in
RUN go build -o /assets/check github.com/dmechas/concourse-team-resource/cmd/check
RUN go build -o /assets/out github.com/dmechas/concourse-team-resource/cmd/out

FROM alpine:edge AS resource
RUN apk add --no-cache bash tzdata ca-certificates
COPY --from=builder assets/ /opt/resource/
RUN chmod +x /opt/resource/*

FROM resource
