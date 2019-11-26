#!/usr/bin/env bash
set -e
echo "" > coverage.txt
# go test -coverprofile=profile.out -covermode=atomic .
# if [ -f profile.out ]; then
#     cat profile.out >> coverage.txt
#     rm profile.out
# fi
#


go test -coverprofile=profile.out -covermode=atomic ./screen
if [ -f profile.out ]; then
    cat profile.out >> coverage.txt
    rm profile.out
fi
go test -coverprofile=profile.out -covermode=atomic ./gesture
if [ -f profile.out ]; then
    cat profile.out >> coverage.txt
    rm profile.out
fi
go test -coverprofile=profile.out -covermode=atomic ./driver/internal/...
if [ -f profile.out ]; then
    cat profile.out >> coverage.txt
    rm profile.out
fi



#go test -coverprofile=profile.out -covermode=atomic ./driver/x11driver/...
#if [ -f profile.out ]; then
#    cat profile.out >> coverage.txt
#    rm profile.out
#fi