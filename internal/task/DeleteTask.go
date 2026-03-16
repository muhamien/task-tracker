package task

import (
	"fmt"
	"syscall"
	"task-tracker/internal/task/domain"
	"task-tracker/pkg"
)

func DeleteTask(id int) error {
	f, err := pkg.LockTaskFile()
	if err != nil {
		return err
	}

	defer syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	defer f.Close()

	tasks, err := pkg.ReadLockedFile(f)
	if err != nil {
		return err
	}

	updatedTasks := make([]domain.Task, 0, len(tasks))
	found := false

	for _, task := range tasks {
		if task.Id == id {
			found = true
			continue
		}
		updatedTasks = append(updatedTasks, task)
	}

	if !found {
		return fmt.Errorf("task Id: %d not found", id)
	}

	fmt.Printf("Task (ID: %d) has been deleted\n", id)
	return pkg.WriteLockedFile(f, updatedTasks)
}
