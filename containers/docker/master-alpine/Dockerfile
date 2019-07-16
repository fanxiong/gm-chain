FROM alpine:3.7

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 --branch release/1.8 https://github.com/fanxiong/gm-chain && \
  (cd gm-chain && make geth) && \
  cp gm-chain/build/bin/mit /mit && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /gm-chain && rm -rf /var/cache/apk/*

EXPOSE 8545
EXPOSE 9999

ENTRYPOINT ["/mit"]
