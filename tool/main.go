package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	// 获取当前目录
	dir, _ := os.Getwd()
	fmt.Println(tree(dir, 0))
}

// go build -o tree.exe main.go
// 目录数生成
// |-- common
// |   |-- config.go
func tree(dir string, level int) string {
	k := "|-- " + readReadmeFile(dir)
	for i := 0; i < level; i++ {
		k = "|    " + k
	}
	// 读取目录
	file, _ := os.ReadDir(dir)
	for _, f := range file {
		// 如果是目录
		if f.IsDir() {
			k += "\n" + tree(path.Join(dir, f.Name()), level+1)
		}
	}
	return k
}

func readReadmeFile(dir string) (str string) {
	dir = path.Join(dir, "README.md")
	// 读取文件,可能会为空
	file, _ := os.ReadFile(dir)
	// 只读取文件第一行
	str = string(file)
	if str == "" {
		str = "empty"
		return
	}

	if len(strings.Split(str, "\n")) < 1 {
		return
	}
	str = strings.Split(str, "\n")[0]
	str = strings.Replace(str, "#", "", -1)
	str = strings.Trim(str, " ")
	return
}
