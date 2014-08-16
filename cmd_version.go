package main

import (
	"flag"
	"fmt"
)

func cmdVersion() command {
	fs := flag.NewFlagSet("nubes version", flag.ExitOnError)

	return command{fs, func(args []string) error {
		fs.Parse(args)
		fmt.Println("nubes")
		fmt.Println("version: " + VERSION)
		return nil
	}}
}
