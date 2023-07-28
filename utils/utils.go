package utils

import (
	"os"
	"unsafe"
)

// DirectoryExists returns whether the given file or directory exists
func DirectoryExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func S2B(s string) (b []byte) {
	return *(*[]byte)(unsafe.Pointer(&s))
}
