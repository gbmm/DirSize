package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var SIZE int64
var MIN_SIZE int

func calcFileSize(path string) int64{
	fi,err:=os.Stat(path)
	var _size int64 = 0
	if err ==nil {
		SIZE += fi.Size()
		_size += fi.Size()
	}
	fmt.Fprintf(os.Stdout,"SIZE=%dM\r", SIZE/1024/1024)
	return _size
}

func parseFile2(path string, ch chan int) {
	files, _ := ioutil.ReadDir(path)
	FileSlice := []string{}
	DirSlice := []string{}
	for _, fi := range files {
		if fi.IsDir() {
			DirSlice = append(DirSlice, path+"/"+fi.Name())
		} else {
			FileSlice = append(FileSlice, path+"/"+fi.Name())
		}
	}

	var count int
	var pathSize int64 =  0
	count = 0
	for _, value := range FileSlice {
		pathSize += calcFileSize(value)
		count++
	}

	if int(pathSize/1024/1024/1024) > MIN_SIZE {
		fmt.Fprintf(os.Stdout,"%s  %dG\n", path,pathSize/1024/1024/1024)
	}

	//find sub directory by go pattern
	DirCount := len(DirSlice)
	if DirCount > 0 {
		DirCH := make([]chan int, DirCount)
		i := 0
		for _, value := range DirSlice {
			DirCH[i] = make(chan int)
			go parseFile2(value, DirCH[i])
			i++
		}

		for _, chs := range DirCH {
			returnCount := <-chs
			count += returnCount
		}
	}
	ch <- count //FindFiles
}


func getFileList(path string, size int) {
	ch := make(chan int)
	MIN_SIZE = size
	go parseFile2(path, ch)
	<-ch
	// fmt.Println("chan<-", <-ch)
	fmt.Println("----------------")
	fmt.Printf("%s TOTAL SIZE=%dG\n",path, SIZE/1024/1024/1024)
}