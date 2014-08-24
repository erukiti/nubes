package main

import (
	"github.com/erukiti/nubes/template"
	// "fmt"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
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

func FetchTemplate(templateName string) (string, error) {
	content, err := ioutil.ReadFile(templateName)
	return string(content), err
}

func Create(name string, templateNames []string) error {
	os.Mkdir(name, 0755)
	for _, templateName := range templateNames {
		var err error
		templateString, err := FetchTemplate(templateName)
		if err != nil {
			return err
		}
		tmpls, err := template.Parse(templateString)
		if err != nil {
			return err
		}
		err = template.Do(tmpls, name)
		if err != nil {
			return err
		}
	}
	return nil
}
