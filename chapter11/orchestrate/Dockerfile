FROM golang:alpine

ENV GOPATH /code/
ADD . /code/src/github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter11/docker
WORKDIR /code/src/github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter11/docker/example
RUN go build

ENTRYPOINT /code/src/github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter11/docker/example/example


