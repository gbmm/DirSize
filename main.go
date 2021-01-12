package main

import (
	"flag"
	"fmt"
	"time"
)

var InputDir = flag.String("dir", "path", "input dir, like C:/temp")
var InputFile = flag.String("file", "n", "output big size file")
var InputSize = flag.Int("size", 1, "input size (G), like 1")

func main()  {
	flag.Parse()
	var path string
	var size int
	var fileParam string
	path = *InputDir
	size = *InputSize
	fileParam = *InputFile
	if path == "path" {
		flag.Usage()
		return
	}
	start := time.Now() // 获取当前时间
	getFileList(path, size, fileParam)
	elapsed := time.Since(start)
	fmt.Println("----------------")
	fmt.Println("\nSPEED TIME：", elapsed)
}



