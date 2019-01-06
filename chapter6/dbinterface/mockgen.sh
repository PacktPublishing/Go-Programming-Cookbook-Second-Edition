#!/usr/bin/env bash
mockgen -destination mocks_test.go -package dbinterface github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter6/dbinterface DB,Transaction
