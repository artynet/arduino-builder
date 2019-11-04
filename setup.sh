#!/bin/bash -x

checkgopath () {

    GOPATH=$(printenv GOPATH)

    if [ -z $GOPATH ]
    then
        mkdir -p ${HOME}/go
        export GOPATH=${HOME}/go
    fi

}

# check if GOPATH variable is blank or not
checkgopath

# detect version
VERSION=`cat main.go| grep "const VERSION" |cut -f4 -d " " | tr -d '"'`

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
