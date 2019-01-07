package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args[0])

	dir, err := os.Open(args[0])
	if err != nil {
		panic(err)
	}
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		panic(err)
	}
	for _, fileInfo := range fileInfos {
		fmt.Printf("%#v\n", fileInfo)
		if fileInfo.IsDir() {
			fmt.Printf("[Dir]  %s\n", fileInfo.Name())
		} else {
			fmt.Printf("[File]  %s\n", fileInfo.Name())
		}
	}
}
