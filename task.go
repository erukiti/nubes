package main

import (
	"./task"
)

func Task(taskName string, taskFile string) error {
	task.Run(taskName, taskFile)
	return nil
}
