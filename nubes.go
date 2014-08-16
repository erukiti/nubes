package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type command struct {
	fs *flag.FlagSet
	fn func(args []string) error
}

func main() {
	commands := map[string]command{
		"version": cmdVersion(),
		"create":  cmdCreate(),
		"task":    cmdTask(),
	}
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: nubes <subcommand> [options]")
		for name, _ := range commands {
			fmt.Fprintf(os.Stderr, "nubes %s\n", name)
		}
		flag.PrintDefaults()
	}

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if cmd, ok := commands[args[0]]; !ok {
		flag.Usage()
		log.Fatalf("Unknown command: %s", args[0])
	} else if err := cmd.fn(args[1:]); err != nil {
		log.Fatal(err)
	}
}

const VERSION = "0.1.0"
