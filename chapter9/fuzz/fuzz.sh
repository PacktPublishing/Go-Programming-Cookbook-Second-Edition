#!/usr/bin/env bash
go-fuzz-build github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter9/fuzz
go-fuzz -bin=./fuzz-fuzz.zip -workdir=output
