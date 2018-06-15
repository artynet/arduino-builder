#!/bin/bash -x

export GOPATH=$PWD

go get github.com/go-errors/errors
go get github.com/stretchr/testify
go get github.com/jstemmer/go-junit-report
go get -u github.com/arduino/go-properties-map
go get -u github.com/arduino/go-timeutils
go get google.golang.org/grpc
go get github.com/golang/protobuf/proto
go get golang.org/x/net/context
go get github.com/fsnotify/fsnotify
go get github.com/arduino/arduino-builder
