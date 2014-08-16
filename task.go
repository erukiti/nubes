package main

import (
	"./task"
	"io/ioutil"
)

func Task(taskName string, taskFile string) error {
	content, err := ioutil.ReadFile(taskFile)
	if err != nil {
		return err
	}

	task.Run(string(content), taskName)
	return nil
}
