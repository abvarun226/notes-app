FROM golang:1.13 AS build-env

RUN apt-get update
RUN apt-get install -y sudo

ENV GOBIN=/go/bin
ENV GOPATH=/go
RUN mkdir -p /go/src /go/bin /go/pkg

ADD . /go/notes-app
WORKDIR /go/notes-app

RUN curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(go env GOPATH)/bin v1.18.0

RUN make build

RUN rm -fr /go/src /go/pkg

# Build an alpine container for the binary.
FROM alpine:latest

RUN wget https://github.com/jwilder/dockerize/releases/download/v0.6.0/dockerize-alpine-linux-amd64-v0.6.0.tar.gz
RUN tar -C /usr/local/bin -xvzf dockerize-alpine-linux-amd64-v0.6.0.tar.gz
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /
COPY --from=build-env /go/bin/note-api /bin/note-api
COPY --from=build-env /go/bin/note-svc /bin/note-svc
CMD ["/bin/note-api;/bin/note-svc"]