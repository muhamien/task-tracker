package task

import (
	"fmt"
	"syscall"
	"task-tracker/internal/task/domain"
	"task-tracker/pkg"
)

func ListTasks(status string) error {
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

	for _, task := range tasks {
		if status == "all" || task.Status == domain.TaskStatus(status) {
			fmt.Printf("ID: %d | Description: %s | Status: %s | CreatedAt: %s | UpdatedAt: %s\n", task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	}

	return nil
}
