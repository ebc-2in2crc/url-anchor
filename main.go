package main

import (
	"os"
)

var version = "0.0.1"

func main() {
	c := cli{outStream: os.Stdout, errStream: os.Stderr}
	c.run(os.Args)
}
