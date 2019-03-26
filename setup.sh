#!/bin/bash -x

checkgopath () {

    GOPATH=$(printenv GOPATH)

    if [ -z $GOPATH ]
    then
        mkdir -p ${HOME}/go
        export GOPATH=${HOME}/go
    fi

}

checkgopath

# cleaning all go packages
rm -rf $GOPATH/{pkg,src}/*

# downloading dependencies
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

cd $GOPATH/src/github.com/arduino/arduino-builder
git checkout 1.4.4
