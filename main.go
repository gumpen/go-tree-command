package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	flag.Parse()
	args := flag.Args()

	rootPath := args[0]
	fmt.Println(rootPath)
	tree(rootPath, 0, "")
}

func tree(filePath string, depth int, ancestralRuledLine string) error {
	dir, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		panic(err)
	}

	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].Name() < fileInfos[j].Name()
	})

	for i, fileInfo := range fileInfos {
		preRuledLine := ""
		if i == len(fileInfos)-1 {
			preRuledLine = "└── "
		} else {
			preRuledLine = "├── "
		}
		fmt.Println(ancestralRuledLine + preRuledLine + fileInfo.Name())
		if fileInfo.IsDir() {
			if i == len(fileInfos)-1 {
				tree(filepath.Join(filePath, fileInfo.Name()), depth+1, ancestralRuledLine+"    ")
			} else {
				tree(filepath.Join(filePath, fileInfo.Name()), depth+1, ancestralRuledLine+"│   ")
			}
		}
	}

	return nil
}
