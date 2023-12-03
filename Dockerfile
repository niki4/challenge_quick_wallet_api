FROM golang:alpine

# copy local package files to the container workspace
ADD . /go/src/github.com/Shyam123-bot/challenge_quick_wallet_api

## install the app
WORKDIR /go/src/github.com/Shyam123-bot/challenge_quick_wallet_api
RUN go build -o /bin/challenge_quick_wallet_api

# start the app
ENTRYPOINT ["/bin/challenge_quick_wallet_api"]
