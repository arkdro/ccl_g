package main

import (
	"github.com/asdf/ccl_g/request"

	"flag"
)

func main() {
	var file = flag.String("infile", "", "input file")
	var dir = flag.String("indir", "", "input file")
	var connectivity = flag.Int("connectivity", 8, "connectivity, 6 or 8")
	var remove = flag.Bool("remove", false, "remove input file on success")
	flag.Parse()
	request.Run(*file, *dir, *remove, *connectivity)
}

