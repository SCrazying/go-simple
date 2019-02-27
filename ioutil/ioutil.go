package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"syscall"
)

type TestUtil struct {
}

type TestUtilInterface interface {
	TestReadFile(filepath string) string
	TestWriteFile(filepath, data string) error
}

func (this TestUtil) TestReadFile(filepath string) string {

	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("ReadFile error")
	}
	return string(bytes)
}

func (this TestUtil) TestWriteFile(filepath, data string) error {
	err := ioutil.WriteFile(filepath, []byte(data), syscall.O_APPEND|os.ModeAppend)
	return err
}

func (this TestUtil) TestReadDir(dirpath string) []string {
	fileinfos, err := ioutil.ReadDir(dirpath)
	if err != nil {
		fmt.Println("ReadDir error")
	}
	var paths = make([]string, 0)
	for _, fileinfo := range fileinfos {
		if fileinfo.IsDir() {
			continue
		}
		paths = append(paths, dirpath+fileinfo.Name())
	}
	return paths
}
func main() {
	var testutil TestUtil
	var data = testutil.TestReadFile("1.txt")
	fmt.Println(data)

	testutil.TestWriteFile("2.txt", data)

	var paths = testutil.TestReadDir("E:/gui/res/ui/")
	fmt.Println(paths)
}
