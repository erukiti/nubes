package main

import (
	"flag"
	"fmt"
	"github.com/erukiti/nubes/task"
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

func Task(taskName string, taskFile string) error {
	task.Run(taskName, taskFile)
	return nil
}
