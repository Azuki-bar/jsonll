package main

import (
	"os"

	"github.com/azuki-bar/jsonll/cli"
)

func main() {
	if err := cli.Main(os.Stdin, os.Stdout, os.Stderr); err != nil {
		panic(err)
	}
}
