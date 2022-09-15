// Copyright 2022 Terin Stock.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/terinjokes/stolport/stol"
)

func main() {
	flag.Parse()

	switch flag.NArg() {
	case 0:
		generate()
	default:
		parse(flag.Arg(0))
	}
}

func generate() {
	// fun, definitely arbitrary epoch
	t := uint16(time.Now().UTC().Unix()/86400 - 7841)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	id, err := stol.New(t, r)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", id)
}

func parse(s string) {
	id, err := stol.Parse(s)
	if err != nil {
		panic(err)
	}

	dc := id.Days()
	e := id.Entropy()
	t := time.Unix(int64(dc+7841)*86400, 0).UTC()

	fmt.Printf("%s:%X\n", t.Format("2006-01-02"), e)
}
