package main

import (
	"flag"
	"fmt"
)

func cmdTask() command {
	fs := flag.NewFlagSet("nubes task", flag.ExitOnError)

	return command{fs, func(args []string) error {
		fs.Parse(args)
		if len(fs.Args()) < 1 {
			return fmt.Errorf("Illegal Argument")
		}

		err := Task(fs.Args()[0], ".nubes.rb")
		return err
	}}
}
