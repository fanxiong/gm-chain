// Copyright 2018 The gm-chain Authors
// This file is part of the gm-chain library.
//
// The gm-chain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The gm-chain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the gm-chain library. If not, see <http://www.gnu.org/licenses/>.

// Contains initialization code for the mbile library.

package mit

import (
	"os"
	"runtime"

	"github.com/fanxiong/gm-chain/log"
)

func init() {
	// Initialize the logger
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stderr, log.TerminalFormat(false))))

	// Initialize the goroutine count
	runtime.GOMAXPROCS(runtime.NumCPU())
}
