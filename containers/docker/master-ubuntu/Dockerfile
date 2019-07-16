FROM ubuntu:xenial

ENV PATH=/usr/lib/go-1.9/bin:$PATH

RUN \
  apt-get update && apt-get upgrade -q -y && \
  apt-get install -y --no-install-recommends golang-1.9 git make gcc libc-dev ca-certificates && \
  git clone --depth 1 --branch release/1.8 https://github.com/fanxiong/gm-chain && \
  (cd gm-chain && make geth) && \
  cp gm-chain/build/bin/mit /mit && \
  apt-get remove -y golang-1.9 git make gcc libc-dev && apt autoremove -y && apt-get clean && \
  rm -rf /gm-chain

EXPOSE 8545
EXPOSE 9999

ENTRYPOINT ["/mit"]
