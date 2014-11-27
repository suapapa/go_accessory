// Copyright 2014, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"io"
	"os"

	"github.com/suapapa/go_accessory"
)

func main() {
	a := accessory.NewContext()
	err := a.SwitchToAccessoryMode(
		"Google, Inc.",
		"AccessoryChat",
		"Accessory Chat",
		"1.0",
		"http://www.android.com",
		"1234567890",
	)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(os.Stdout, a)
	if err != nil {
		panic(err)
	}

	// _, err = io.Copy(a, os.Stdin)
	// if err != nil {
	// 	panic(err)
	// }
}
