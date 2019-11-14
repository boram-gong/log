package log

import (
	"fmt"
	"os"
)

func IsDir(name string) bool {
	fi, err := os.Stat(name)
	if err != nil {
		return false
	}

	return fi.IsDir()
}

func CreateDir(name string) {
	if IsDir(name) {
		return
	}
	if createDirImpl(name) {
		fmt.Println("create dir successfully:", name)
		return
	}
}

func createDirImpl(name string) bool {
	err := os.MkdirAll(name, os.ModePerm)
	if err == nil {
		return true
	} else {
		fmt.Println("create dir Error: ", err)
		return false
	}
}
