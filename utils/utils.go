package utils

import (
	"io"
	"log"
	"os"
)

func IsDir(file string) bool {
	fileInfo, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}
	return fileInfo.IsDir()
}

func IsFile(file string) bool {
	fileInfo, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}
	return !fileInfo.IsDir()
}

func MoveFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return os.Remove(src)
}
