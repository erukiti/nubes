package main

import (
	"flag"
	"fmt"
)

func cmdCreate() command {
	fs := flag.NewFlagSet("nubes create", flag.ExitOnError)

	return command{fs, func(args []string) error {
		fs.Parse(args)
		if len(fs.Args()) < 2 {
			return fmt.Errorf("Illegal Argument")
		}

		err := Create(fs.Args()[0], fs.Args()[1:])
		return err
	}}
}
