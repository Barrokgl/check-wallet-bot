FROM golang

ADD . /go/src/github.com/Barrokgl/check-wallet-bot

RUN go install github.com/Barrokgl/check-wallet-bot

ENTRYPOINT /go/bin/check-wallet-bot

EXPOSE 5005
