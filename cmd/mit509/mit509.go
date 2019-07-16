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

// mit is the official command-line client for Mit.

package main

import (
	"github.com/fanxiong/gm-chain/mitx509"
	"fmt"
)

var (
	country = []string{"CN"}
	orgnazation = []string{"weyii.co"};
	orgnazationalUnit = []string{"weyii.co"}
)

func main(){
	x509_issuer := &mitx509.X509mit{}
	x509_issuer = x509_issuer.GenCa(country,orgnazation,orgnazationalUnit)
	fmt.Printf("%v",x509_issuer)
	return;
}
