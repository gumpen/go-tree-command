package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type flags struct {
	a *bool
	d *bool
}

func main() {
	var flags flags
	flags.a = flag.Bool("a", false, "show dot prefix files")
	flags.d = flag.Bool("d", false, "show dir only")
	flag.Parse()
	args := flag.Args()

	rootPath := args[0]
	fmt.Println(rootPath)
	tree(rootPath, 0, "", flags)
}

func tree(filePath string, depth int, ancestralRuledLine string, flags flags) error {
	dir, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		panic(err)
	}

	if !(*flags.a) {
		fileInfosExcluded := make([]os.FileInfo, 0, len(fileInfos))
		for _, fileInfo := range fileInfos {
			if !strings.HasPrefix(fileInfo.Name(), ".") {
				fileInfosExcluded = append(fileInfosExcluded, fileInfo)
			}
		}
		fileInfos = fileInfosExcluded
	}

	if *flags.d {
		fileInfosOnlyDir := make([]os.FileInfo, 0, len(fileInfos))
		for _, fileInfo := range fileInfos {
			if fileInfo.IsDir() {
				fileInfosOnlyDir = append(fileInfosOnlyDir, fileInfo)
			}
		}
		fileInfos = fileInfosOnlyDir
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
				tree(filepath.Join(filePath, fileInfo.Name()), depth+1, ancestralRuledLine+"    ", flags)
			} else {
				tree(filepath.Join(filePath, fileInfo.Name()), depth+1, ancestralRuledLine+"│   ", flags)
			}
		}
	}

	return nil
}
