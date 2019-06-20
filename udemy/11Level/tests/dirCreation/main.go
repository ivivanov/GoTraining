package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	createDirectories([]string{"one", "two"})
}

func createDirectories(directories []string) {
	for _, dir := range directories {
		err := os.Mkdir(fmt.Sprintf("%v/%v", getRootPath(), dir), os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func getRootPath() string {
	wd, _ := os.Getwd()
	dir := path.Dir(wd)
	root := path.Join(dir, "../")
	return root
}
