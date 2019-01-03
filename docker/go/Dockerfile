FROM golang:1.10

RUN apt-get update -qq && \
    apt-get install -y mysql-client vim && \
    go get -u github.com/go-sql-driver/mysql && \
    go get -u github.com/prometheus/client_golang/prometheus/promhttp

WORKDIR /go/src/sample_docker_compose

ADD . .

RUN go build
