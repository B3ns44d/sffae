package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		originalFilePath := path.Join(pwd, file.Name())
		if isFile(originalFilePath) && !strings.HasPrefix(file.Name(), ".") {
			fileName := path.Base(originalFilePath)
			ext := path.Ext(fileName)
			extensionPath := path.Join(pwd, ext)
			if _, err := os.Stat(extensionPath); os.IsNotExist(err) {
				os.Mkdir(extensionPath, os.ModePerm)
			}
			err := CopyFile(originalFilePath, path.Join(extensionPath, fileName))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func isFile(file string) bool {
	fileInfo, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}
	return !fileInfo.IsDir()
}

func CopyFile(src, dst string) error {
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
	return out.Close()
}
