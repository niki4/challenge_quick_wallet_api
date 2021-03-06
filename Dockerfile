FROM golang:alpine

ADD . /go/src/github.com/niki4/challenge_quick_wallet_api
RUN go install github.com/niki4/challenge_quick_wallet_api@latest
