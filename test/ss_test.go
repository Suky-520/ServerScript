package test

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"ss"
	"strings"
	"testing"
)

var testPath = map[string]string{
	"index": "./index.js",
	"test":  "./ss/for/while.js",
}

// 测试index
func Test_index(t *testing.T) {
	v := ss.NewVM()
	v.RunJsFile(testPath["index"])
}

// 测试编译index
func Test_compileIndex(t *testing.T) {
	v := ss.NewVM()
	path := testPath["test"]
	ssPath := ""
	extIndex := strings.LastIndex(path, ".")
	if extIndex > -1 {
		ssPath = path[:extIndex] + ".ss"
	}
	v.Build(path, ssPath)
	v.RunBinFile(ssPath)
	// 删除生成的ss文件
	os.Remove(ssPath)
}

// 测试所有用例
func Test_all(t *testing.T) {
	i := 0
	for _, p := range getFiles("./") {
		fmt.Println(p)
		v := ss.NewVM()
		v.RunJsFile("./" + p)
		i++
	}
	fmt.Println("共完成", i, "个测试用例")
}

// 测试编译所有用例
func Test_compileAll(t *testing.T) {
	i := 0
	for _, path := range getFiles("./") {
		fmt.Println(path)
		v := ss.NewVM()
		// 获取不包含扩展名的文件名部分
		extIndex := strings.LastIndex(path, ".")
		ssPath := ""
		if extIndex > -1 {
			ssPath = path[:extIndex] + ".ss"
		}
		//fmt.Println(ssPath)
		v.Build("./"+path, "./"+ssPath)
		v.RunBinFile("./" + ssPath)
		// 删除生成的ss文件
		os.Remove("./" + ssPath)
		i++
	}
	fmt.Println("共完成", i, "个测试用例")
}

func getFiles(dir string) []string {
	// 读取目录中的所有文件和子目录
	list := make([]string, 0)
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, "js") && path != "index.js" {
			list = append(list, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking through directory:", err)
	}
	return list
}
