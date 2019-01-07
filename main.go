// package main

// import (
// 	"flag"
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"strings"
// )

// func main() {
// 	flag.Parse()
// 	args := flag.Args()
// 	fmt.Println(args[0])
// 	dir := filepath.Dir(args[0])
// 	fmt.Printf("%#v\n", dir)

// 	// tree(args[0], 0)
// 	// dir, err := os.Open(args[0])
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// fileInfos, err := dir.Readdir(-1)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// fmt.Println(".")
// 	// for i, fileInfo := range fileInfos {
// 	// 	pre_ruled_line := ""
// 	// 	if i == len(fileInfos)-1 {
// 	// 		pre_ruled_line = "└"
// 	// 	} else {
// 	// 		pre_ruled_line = "├"
// 	// 	}
// 	// 	fmt.Println(pre_ruled_line + fileInfo.Name())
// 	// }
// }

// func tree(filepath string, depth int) error {
// 	// 再帰したい
// 	dir, err := os.Open(filepath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fileInfos, err := dir.Readdir(-1)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for i, fileInfo := range fileInfos {
// 		preRuledLine := ""
// 		if i == len(fileInfos)-1 {
// 			preRuledLine = "└"
// 		} else {
// 			preRuledLine = "├"
// 		}
// 		fmt.Println(strings.Repeat("|", depth) + preRuledLine + fileInfo.Name())
// 		if fileInfo.IsDir() {
// 			tree(fileInfo, depth+1)
// 		}
// 	}

// 	return nil
// }

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func main() {

	// 再帰なし
	files, _ := filepath.Glob("./*")
	for _, f := range files {
		printPathAndSize(f)
	}

	fmt.Println("---------")

	// 再帰あり
	filepath.Walk("./example/", visit)

}

func visit(path string, info os.FileInfo, err error) error {

	printPathAndSize(path)
	return nil
}

func printPathAndSize(path string) {
	// ファイルサイズの取得
	var s syscall.Stat_t
	syscall.Stat(path, &s)

	fmt.Print(path)
	fmt.Print(": ")
	fmt.Print(s.Size)
	fmt.Println(" bytes")

}
