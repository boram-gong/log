package log

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
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

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

func RmFile(path string, name string) {
	fileInfoList, _ := ioutil.ReadDir(path)
	var rmList []string
	for i := range fileInfoList {
		if ok, _ := regexp.MatchString(name, fileInfoList[i].Name()); ok {
			rmList = append(rmList, fileInfoList[i].Name())
		}
	}
	for _, r := range rmList {
		os.Remove(path + "/" + r)
	}
}
