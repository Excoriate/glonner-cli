package system

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

type DirReadResult struct {
	Files []string
	Dirs  []string
}

func CheckIfDirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func ReadDir(path string) (DirReadResult, error) {
	var files []string
	var dirs []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		} else {
			dirs = append(dirs, path)
		}
		return nil
	})

	if err != nil {
		return DirReadResult{}, err
	}

	return DirReadResult{Files: files}, nil
}

func DeleteDir(path string, deleteAllContent bool) (bool, error) {
	if !CheckIfDirExists(path) {
		return false, errors.New(fmt.Sprintf("Directory %s does not exist", path))
	}

	if deleteAllContent {
		err := os.RemoveAll(path)
		if err != nil {
			return false, err
		}
	} else {
		err := os.Remove(path)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func CheckIfDirHasContent(path string, ignoredFile []string) bool {
	files, err := ReadDir(path)

	if err != nil {
		return false
	}

	if len(files.Files) != 0 || len(files.Dirs) != 0 {
		if len(ignoredFile) != 0 {
			var filesToCheck []string

			for _, file := range files.Files {
				for _, ignored := range ignoredFile {
					justTheFileName := filepath.Base(file)
					if justTheFileName == ignored {
						filesToCheck = append(filesToCheck, file)
					}
				}
			}

			if len(filesToCheck) == len(files.Files) {
				return false
			}
		}

		return true
	}

	return false
}

func GetHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	return home
}

func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
