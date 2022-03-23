package main

import (
	"os"

	"github.com/waelsy123/invasion"
)

func main() {
	exitCode := invasion.Main(os.Args)
	os.Exit(exitCode)
}
