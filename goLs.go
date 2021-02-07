package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// MyFile is like FileInfo but simpler
type MyFile struct {
	name  string
	size  int64
	isDir bool
}

func (f MyFile) String() string {
	return fmt.Sprintf("Name:%s Size:%d Dir:%t\n", f.name, f.size, f.isDir)
}

func main() {
	fileinfo, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	sliceMyFiles := make([]MyFile, 0, len(fileinfo))
	for _, line := range fileinfo {
		mf := MyFile{
			name:  line.Name(),
			size:  line.Size(),
			isDir: line.IsDir(),
		}
		sliceMyFiles = append(sliceMyFiles, mf)
	}

	for _, slice := range sliceMyFiles {
		fmt.Println(slice)
	}
}
