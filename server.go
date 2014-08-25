package main

import (
	"flag"
	"github.com/erukiti/nubes/task"
	// "os/signal"
	// "syscall"
)

func cmdServer() command {
	fs := flag.NewFlagSet("nubes task", flag.ExitOnError)

	return command{fs, func(args []string) error {
		var defaultScript string
		fs.StringVar(&defaultScript, "script", "nubes/task.rb", "task script")
		fs.Parse(args)

		// sig := make(chan os.Signal)
		// signal.Notify(sig, syscall.SIGHUP)
		t := task.New()
		t.Load(defaultScript)

		for {
			select {}
		}
	}}
}
