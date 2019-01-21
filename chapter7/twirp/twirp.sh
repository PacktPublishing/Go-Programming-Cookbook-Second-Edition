#!/usr/bin/env bash
protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./rpc/greeter/greeter.proto
