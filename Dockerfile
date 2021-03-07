FROM golang:1.16rc1-buster
WORKDIR /go/src/work
ENV GO111MODULE=on
# CMD ["go", "run", "/go/src/main.go"]