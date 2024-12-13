package main

import (
    "crypto/sha256"
    "fmt"
    "io"
    "os"
    "time"
)

func generateChecksum(file string) (string, error) {
    startTime := time.Now()

    hashAlgorithm := sha256.New()

    f, err := os.Open(file)
    if err != nil {
        return "", fmt.Errorf("file not found")
    }
    defer f.Close()

    buffer := make([]byte, 8192)
    for {
        n, err := f.Read(buffer)
        if err != nil && err != io.EOF {
            return "", err
        }
        if n == 0 {
            break
        }
        hashAlgorithm.Write(buffer[:n])
    }

    duration := time.Since(startTime).Milliseconds()
    fmt.Println(duration)

    return fmt.Sprintf("%x", hashAlgorithm.Sum(nil)), nil
}

func main() {
    checksum, err := generateChecksum("test_file.dpx")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(checksum)
}