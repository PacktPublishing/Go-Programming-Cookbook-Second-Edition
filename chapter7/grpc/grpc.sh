#!/usr/bin/env bash
protoc --go_out=plugins=grpc:. greeter/greeter.proto  