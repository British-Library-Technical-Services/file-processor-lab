package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type FileData struct {
	File *os.File
}

func ReadFile(file string) (*FileData, error) {
	data, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return &FileData{File: data}, nil
}

func (fd *FileData) GenerateChecksum() (string, error) {
	hashAlgorithm := sha256.New()
	buffer := make([]byte, 8192)

	for {
		n, err := fd.File.Read(buffer)
		if err != nil && err != io.EOF {
			return "", err
		}
		if n == 0 {
			break
		}
		hashAlgorithm.Write(buffer[:n])
	}
	defer fd.File.Close()

	hashByte := hashAlgorithm.Sum(nil)
	hashHex := hex.EncodeToString(hashByte)

	return hashHex, nil
}

func main() {

	start := time.Now()

	dirPath := "/workspaces/file-processor-lab/_test_files"
	filePath := filepath.Join(dirPath, "*.m4a")

	fileList, err := filepath.Glob(filePath)
	if err != nil {
		fmt.Println(err)
	}

	var wg sync.WaitGroup

	for _, file := range fileList {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()

			fd, err := ReadFile(file)
			if err != nil {
				fmt.Println(err)
			}

			checksum, err := fd.GenerateChecksum()
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(checksum)
		}(file)
	}

	wg.Wait()

	duration := time.Since(start)
	fmt.Println("Time taken: ", duration)
}
