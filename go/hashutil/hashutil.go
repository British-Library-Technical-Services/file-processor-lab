package hashutil

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"file-processor-lab/fileutil"
)

func GenerateChecksum(fd *fileutil.FileData) (string, error) {
	algorithm := sha256.New()
	buffer := make([]byte, 8192)

	for {
		n, err := fd.File.Read(buffer)
		if err != nil && err != io.EOF {
			return "", err
		}
		if n == 0 {
			break
		}
		algorithm.Write(buffer[:n])
	}
	defer fd.File.Close()

	inBytes := algorithm.Sum(nil)
	inHex := hex.EncodeToString(inBytes)

	return inHex, nil
}