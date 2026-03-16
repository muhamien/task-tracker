package task

import (
	"fmt"
	"syscall"
	"task-tracker/internal/task/domain"
	"task-tracker/pkg"
	"time"
)

func UpdateTask(id int, description string, status ...string) error {
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

	taskExist := false
	var updatedTask []domain.Task
	for _, task := range tasks {
		if task.Id == id {
			taskExist = true
			if len(description) > 0 {
				task.Description = description
			}
			if len(status) > 0 {
				task.Status = domain.TaskStatus(status[0])
			}
			task.UpdatedAt = time.Now()
		}
		updatedTask = append(updatedTask, task)
	}

	if !taskExist {
		return fmt.Errorf("task Id: %d not found", id)
	}

	fmt.Printf("Task updated successfully: (ID: %d)", id)
	return pkg.WriteLockedFile(f, updatedTask)
}
