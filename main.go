package main

import (
	"fmt"
	"os"
)

func arg(args []string) string {
	if len(args) > 1 {
		return args[1]
	}
	return ""
}

func exitWithError(err error) {
	fmt.Printf("ERROR: %v\n", err)
	os.Exit(1)
}

func main() {
	newGame(arg(os.Args), os.Stdin, os.Stdout).start()
}
