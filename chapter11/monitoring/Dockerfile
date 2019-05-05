FROM golang:alpine

ENV GOPATH /code/
ADD . /code/src/github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter11/monitoring
WORKDIR /code/src/github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter11/monitoring
RUN go build

ENTRYPOINT /code/src/github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter11/monitoring/monitoring
