# start a golang base image, version 1.8
FROM golang:1.9

#switch to our app directory
COPY . /go/src/github.com/Ivandolchevic/goapis/  
WORKDIR /go/src/github.com/Ivandolchevic/goapis/cmd/draft

#disable crosscompiling 
ENV CGO_ENABLED=0

#compile linux only
ENV GOOS=linux

RUN go get ./

#build the binary with debug information removed
RUN go build  -ldflags '-w -s' -a -installsuffix cgo -o draft 
