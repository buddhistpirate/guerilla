package main

import (
	"fmt"
	"guerillaradio"
	"os"
)

func main() {
	library := guerillaradio.Library{}
	err := library.AddDirectory("fixtures")
	if err != nil {
		fmt.Println("Error opening fixtures")
		os.Exit(1)
	}
	guerillaradio.Listen(&library, ":8080)
}
