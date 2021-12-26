package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/B3ns44d/sffae/utils"
)

func main() {
	pwd := handlePath(os.Args[1:])

	files, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}
	var count int
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		count++
	}

	if count == 0 {
		fmt.Println("No files in this directory")
		fmt.Println("Exiting...")
		os.Exit(0)
	}

	fmt.Printf("Do you want to move all %d files in this directory? (Y/n): ", count)
	var answer string
	fmt.Scanln(&answer)
	if answer == "n" || answer == "N" {
		fmt.Println("Exiting...")
		os.Exit(0)
	}
	for _, file := range files {
		originalFilePath := path.Join(pwd, file.Name())
		if utils.IsFile(originalFilePath) && !strings.HasPrefix(file.Name(), ".") {
			fileName := path.Base(originalFilePath)
			ext := strings.Replace(path.Ext(fileName), ".", "", -1)
			extensionPath := path.Join(pwd, ext)
			fmt.Printf("Moving %s to %s\n", fileName, extensionPath)
			if _, err := os.Stat(extensionPath); os.IsNotExist(err) {
				fmt.Println("Creating directory:", extensionPath)
				os.Mkdir(extensionPath, os.ModePerm)
			}
			err := utils.MoveFile(originalFilePath, path.Join(extensionPath, fileName))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	fmt.Println("Done!")
}

func handlePath(args []string) string {
	if len(args) == 0 {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		return pwd
	}
	if utils.IsFile(args[0]) {
		return path.Dir(args[0])
	}
	if utils.IsDir(args[0]) {
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
