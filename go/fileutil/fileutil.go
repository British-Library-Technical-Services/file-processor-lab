package fileutil

import (
	"os"
	"path/filepath"
)

type FileData struct {
	File *os.File
}

func ReadFile(file string) (*FileData, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return &FileData{File: f}, nil
}

func FileCleanUp() (string, error) {
	dirPath := "/workspaces/file-processor-lab/_test_files"
	filePath := filepath.Join(dirPath, "*.mp3")

	cleanUpList, err := filepath.Glob(filePath)
	if err != nil {
		return "", err
	}
	
	for _, file := range cleanUpList {
		err := os.Remove(file)
		if err != nil {
			return "", err
		}
	}
	return "Files cleaned up", nil
}