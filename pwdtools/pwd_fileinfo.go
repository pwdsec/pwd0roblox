package pwdtools

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

// get md5 hash of file
func GetHash(file string) string {
	return getHash(file)
}

func GetSize(file string) int64 {
	return getSize(file)
}

func getHash(file string) string {
	f, err := os.Open(file)
	if err != nil {
		return ""
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// get size of file in mb
func getSize(file string) int64 {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return 0
	}
	return fileInfo.Size() / 1024 / 1024
}
