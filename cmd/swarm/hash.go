// Copyright 2018 The gm-chain Authors
// This file is part of gm-chain.
//
// gm-chain is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// gm-chain is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with gm-chain. If not, see <http://www.gnu.org/licenses/>.

// Command bzzhash computes a swarm tree hash.
package main

import (
	"fmt"
	"os"

	"github.com/fanxiong/gm-chain/cmd/utils"
	"github.com/fanxiong/gm-chain/swarm/storage"
	"gopkg.in/urfave/cli.v1"
)

func hash(ctx *cli.Context) {
	args := ctx.Args()
	if len(args) < 1 {
		utils.Fatalf("Usage: swarm hash <file name>")
	}
	f, err := os.Open(args[0])
	if err != nil {
		utils.Fatalf("Error opening file " + args[1])
	}
	defer f.Close()

	stat, _ := f.Stat()
	chunker := storage.NewTreeChunker(storage.NewChunkerParams())
	key, err := chunker.Split(f, stat.Size(), nil, nil, nil)
	if err != nil {
		utils.Fatalf("%v\n", err)
	} else {
		fmt.Printf("%v\n", key)
	}
}
