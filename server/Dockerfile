FROM golang:1.22-alpine3.19 AS builder

RUN apk add git
WORKDIR /go/src
RUN git clone --depth 1 --branch v0.1.5 https://github.com/shadowsocks/go-shadowsocks2
WORKDIR /go/src/go-shadowsocks2
RUN go build

FROM alpine:3.19 AS distr

COPY --from=builder /go/src/go-shadowsocks2/go-shadowsocks2 /usr/bin/go-shadowsocks2

ENTRYPOINT ["go-shadowsocks2"]