package _file

import (
	"os"
	"strings"
)

const PARENT_DIRECTORY  = "/Users/dante/Desktop/"

func Mkdir(path, cascadePath string) {
	var path1 = strings.Join([]string{PARENT_DIRECTORY, path}, "")
	os.Mkdir(path1, os.ModePerm)
	os.Remove(path1)

	var path2 = strings.Join([]string{PARENT_DIRECTORY, cascadePath}, "")
	os.MkdirAll(path2, os.ModePerm)
	os.RemoveAll(path2)
}


