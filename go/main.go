package main

import (
	"fmt"
	"path/filepath"
	"sync"
	"time"
	"file-processor-lab/fileutil"
	"file-processor-lab/hashutil"
	"file-processor-lab/transcodeutil"
)

func main() {

	start := time.Now()

	dirPath := "/workspaces/file-processor-lab/_test_files"
	filePath := filepath.Join(dirPath, "*.wav")

	fileList, err := filepath.Glob(filePath)
	if err != nil {
		fmt.Println(err)
	}

	var wg sync.WaitGroup

	for _, file := range fileList {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()

			fd, err := fileutil.ReadFile(file)
			if err != nil {
				fmt.Println(err)
			}

			checksum, err := hashutil.GenerateChecksum(fd)
			if err != nil {
				fmt.Println(err)
			}


			transcode, err := transcodeutil.FileTranscode(fd)
			if err != nil {
				fmt.Print(err)
			}

			fmt.Println(checksum)
			fmt.Println(transcode)

		}(file)
	}

	wg.Wait()

	duration := time.Since(start)
	fmt.Println("Time taken: ", duration)

	_, err = fileutil.FileCleanUp() 
	if err != nil {
		fmt.Println(err)
	}
	
}
