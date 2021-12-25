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
	pwd := handlePath(os.Args[1:])

	files, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		originalFilePath := path.Join(pwd, file.Name())
		if isFile(originalFilePath) && !strings.HasPrefix(file.Name(), ".") {
			fileName := path.Base(originalFilePath)
			ext := path.Ext(fileName)
			extensionPath := strings.Replace(path.Join(pwd, ext), ".", "", 1)
			if _, err := os.Stat(extensionPath); os.IsNotExist(err) {
				os.Mkdir(extensionPath, os.ModePerm)
			}
			err := moveFile(originalFilePath, path.Join(extensionPath, fileName))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func handlePath(args []string) string {
	if len(args) == 0 {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		return pwd
	}
	if isFile(args[0]) {
		return path.Dir(args[0])
	}
	if isDir(args[0]) {
		if args[0] == "." {
			pwd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			return pwd
		}
		return args[0]
	}
	return ""
}

func isDir(file string) bool {
	fileInfo, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}
	return fileInfo.IsDir()
}

func isFile(file string) bool {
	fileInfo, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}
	return !fileInfo.IsDir()
}

func moveFile(src, dst string) error {
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
