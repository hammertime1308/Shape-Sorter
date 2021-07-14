package main

import (
	"flag"

	"github.com/hammertime1308/shape-sorter/cli"
	"github.com/hammertime1308/shape-sorter/file"
	"github.com/hammertime1308/shape-sorter/server"
)

func main() {
	// create flags for json, json-file and port
	jsonPtr := flag.String("json", "", "raw json string as input")
	jsonFilePtr := flag.String("json-file", "", "json file as input path")
	portPtr := flag.Int("port", 8080, "Start web service")

	flag.Parse()
	// by default starts the web service if no flags are provided
	if *jsonPtr != "" {
		// parses from cli
		cli.Solve(*jsonPtr)
	} else if *jsonFilePtr != "" {
		// parse from file
		file.Solve(*jsonFilePtr)
	} else {
		// start web service
		server.Serve(*portPtr)
	}
}
