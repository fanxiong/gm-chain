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

package mitclient

import "github.com/fanxiong/gm-chain"

// Verify that Client implements the fanxiong interfaces.
var (
	_ = fanxiong.ChainReader(&Client{})
	_ = fanxiong.TransactionReader(&Client{})
	_ = fanxiong.ChainStateReader(&Client{})
	_ = fanxiong.ChainSyncReader(&Client{})
	_ = fanxiong.ContractCaller(&Client{})
	_ = fanxiong.GasEstimator(&Client{})
	_ = fanxiong.GasPricer(&Client{})
	_ = fanxiong.LogFilterer(&Client{})
	_ = fanxiong.PendingStateReader(&Client{})
	// _ = fanxiong.PendingStateEventer(&Client{})
	_ = fanxiong.PendingContractCaller(&Client{})
)
