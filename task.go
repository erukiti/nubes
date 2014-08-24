package main

import (
	"github.com/erukiti/nubes/task"
)

func Task(taskName string, taskFile string) error {
	task.Run(taskName, taskFile)
	return nil
}
