FROM golang:latest

WORKDIR /go/src/app

RUN go install github.com/cweill/gotests/...@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go install github.com/fatih/gomodifytags@latest
RUN go install github.com/josharian/impl@latest
RUN go install github.com/haya14busa/goplay/cmd/goplay@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install golang.org/x/tools/gopls@latest