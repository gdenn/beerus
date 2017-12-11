FROM golang

ENV GOBIN $GOPATH/bin

ADD . /go/src/github.com/beerus/manager
RUN go install /go/src/github.com/beerus/manager/cmd/api/main.go
WORKDIR /go/src/github.com/beerus/manager
ENTRYPOINT /go/bin/main