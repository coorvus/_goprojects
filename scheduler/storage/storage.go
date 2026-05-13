package storage

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"

	"strconv"
	"syscall"
)

type Task interface {
	Save()
}

func loadFile(filepath string, flag int) (*os.File, error) {
	f, err := os.OpenFile(filepath, flag, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	// Exclusive lock obtained on the file descriptor
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}

func closeFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}

func Save(t []string) error {

	tasks, err := GetAll()
	if err != nil {
		return err
	}

	f, err := loadFile("./storage/files/data.csv", os.O_APPEND|os.O_WRONLY)
	if err != nil {
		return err
	}

	defer closeFile(f)

	w := csv.NewWriter(f)
	size := len(tasks)

	if size == 0 {
		// Write header
		ct := [][]string{
			{"ID", "Description", "CreatedAt", "IsComplete"},
			t,
		}

		if err := w.WriteAll(ct); err != nil {
			return err
		}

	} else {
		t[0] = strconv.FormatInt(int64(size), 10)

		if err := w.Write(t); err != nil {
			return err
		}

		w.Flush()
	}
	if err := w.Error(); err != nil {
		return err
	}

	return nil
}

func GetAll() ([][]string, error) {
	f, err := loadFile("./storage/files/data.csv", os.O_RDONLY|os.O_CREATE)
	if err != nil {
		return [][]string{}, err
	}

	defer closeFile(f)

	r := csv.NewReader(f)
	r.Comma = ','
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		if err == io.EOF {
			return [][]string{}, nil
		}
		return [][]string{}, err
	}

	return records, nil
}

func Complete(idx int) error {
	tasks, err := GetAll()
	if err != nil {
		return err
	}
	f, err := loadFile("./storage/files/data.csv", os.O_WRONLY|os.O_TRUNC)
	if err != nil {
		return err
	}

	defer closeFile(f)

	size := len(tasks)

	if idx == 0 || idx >= size {
		return errors.New("Invalid ID, can not proceed!")
	}

	foundTask := tasks[idx]

	if foundTask[3] == "true" {
		return fmt.Errorf("Todo `%s` was already completed", foundTask[1])
	}
	foundTask[3] = "true"

	tasks[idx] = foundTask

	w := csv.NewWriter(f)

	if err = w.WriteAll(tasks); err != nil {
		return err
	}

	return nil
}

func Delete(idx int) error {
	tasks, err := GetAll()
	if err != nil {
		return err
	}
	f, err := loadFile("./storage/files/data.csv", os.O_WRONLY|os.O_TRUNC)
	if err != nil {
		return err
	}

	defer closeFile(f)

	size := len(tasks)

	if idx == 0 || idx >= size {
		return errors.New("Invalid ID, can not proceed!")
	}

	tasks = slices.Delete(tasks, idx, idx+1)

	w := csv.NewWriter(f)

	if err = w.WriteAll(tasks); err != nil {
		return err
	}

	return nil
}
