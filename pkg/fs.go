package pkg

import (
	"encoding/json"
	"os"
	"path"
	"syscall"
	"task-tracker/internal/task/domain"
)

func GetFilePath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path.Join(dir, "tasks.json")
}

// LockTaskFile opens the tasks file and acquires an exclusive lock.
// The caller must call Unlock() and Close() on the returned file when done.
func LockTaskFile() (*os.File, error) {
	filePath := GetFilePath()
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	if err := syscall.Flock(int(file.Fd()), syscall.LOCK_EX); err != nil {
		file.Close()
		return nil, err
	}
	return file, nil
}

func ReadLockedFile(f *os.File) ([]domain.Task, error) {
	f.Seek(0, 0)
	tasks := []domain.Task{}

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if stat.Size() == 0 {
		return tasks, nil
	}

	if err := json.NewDecoder(f).Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func WriteLockedFile(f *os.File, tasks []domain.Task) error {
	f.Truncate(0)
	f.Seek(0, 0)
	return json.NewEncoder(f).Encode(tasks)
}
