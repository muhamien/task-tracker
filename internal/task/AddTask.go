package task

import (
	"fmt"
	"syscall"
	"task-tracker/internal/task/domain"
	"task-tracker/pkg"
)

func AddTask(description string) error {
	f, err := pkg.LockTaskFile()
	if err != nil {
		return err
	}
	defer f.Close()
	defer syscall.Flock(int(f.Fd()), syscall.LOCK_UN)

	tasks, err := pkg.ReadLockedFile(f)
	if err != nil {
		return err
	}

	var newTaskId int
	if len(tasks) > 0 {
		lastTask := tasks[len(tasks)-1]
		newTaskId = lastTask.Id + 1
	} else {
		newTaskId = 1
	}

	task := domain.NewTask(newTaskId, description)
	tasks = append(tasks, *task)
	
	fmt.Printf("Task added successfully (Id: %d)", task.Id)
	return pkg.WriteLockedFile(f, tasks)
}
