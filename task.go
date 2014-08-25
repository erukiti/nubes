package main

import (
	"flag"
	"fmt"
	"github.com/erukiti/nubes/task"
)

func cmdTask() command {
	fs := flag.NewFlagSet("nubes task", flag.ExitOnError)

	return command{fs, func(args []string) error {
		var defaultScript string
		fs.StringVar(&defaultScript, "script", ".nubes/task.rb", "task script")
		fs.Parse(args)
		if len(fs.Args()) < 1 {
			return fmt.Errorf("Illegal Argument")
		}

		t := task.New()
		t.Load(defaultScript)
		t.RunString(fs.Args()[0])

		return nil
	}}
}
