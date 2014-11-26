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

	go io.Copy(os.Stdout, a)
	io.Copy(a, os.Stdin)
}
