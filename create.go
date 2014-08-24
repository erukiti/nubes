package main

import (
	"github.com/erukiti/nubes/template"
	// "fmt"
	"io/ioutil"
	"os"
)

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
