.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/fanxiong/gm-chain/cmd/swarm
	go build -o bin/mit     github.com/fanxiong/gm-chain/cmd/gmit
	go build -o bin/bootnode github.com/fanxiong/gm-chain/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/fanxiong/gm-chain/swarm/...
