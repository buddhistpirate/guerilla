package main

import (
	"flag"
	"fmt"
	"github.com/buddhistpirate/guerillaRadio"
	"os"
)

func main() {

	directory := flag.String("directory", "fixtures", "Directory of text files to load")
	network_interface := flag.String("interface", ":8080", "Network Interface:Port to listen on")
	flag.Parse()

	library := guerillaradio.Library{}
	err := library.AddDirectory(*directory)
	if err != nil || library.Size() == 0 {
		fmt.Printf("Error loading %v, %v files loaded: %v\n", *directory, library.Size(), err)
		os.Exit(1)
	}
	fmt.Printf("Loaded %v Files from %v\n", library.Size(), *directory)

	guerillaradio.Listen(&library, *network_interface)
}
