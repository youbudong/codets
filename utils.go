package main

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

func GetFuncName(path string) string {
	// and removing the leading slash.
	// nameReplaceFrist := strings.Replace(path, "{", "by/", -1)
	nameReplace := strings.ReplaceAll(path, "{", "")
	nameReplace = strings.ReplaceAll(nameReplace, "}", "")
	nameSplit := strings.Split(nameReplace, "/")
	nameStr := ""
	for _, name := range nameSplit {
		nameStr += firstUpper(name)
	}
	return nameStr
}
func GetOsFile(dir, filename string) *os.File {
	pwd := GetDir(dir)
	filepath := filepath.Join(pwd, filename)
	// 文件存在则覆盖 否则创建
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	return file
}

func GetDir(d string) string {
	curDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir := filepath.Join(curDir, d)
	// 确保目录存在
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic(err)
	}
	return dir
}

func GetFilepath(path string) string {
	// 获取当前工作目录
	curDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// 拼接路径
	filepath := filepath.Join(curDir, path)
	return filepath
}

// 首字母大写
func firstUpper(str string) string {
	if len(str) == 0 {
		return ""
	}
	return strings.ToUpper(str[0:1]) + str[1:]
}

// sortMapKeys 接受一个 map 和一个函数，该函数用于处理排序后键值对。
func sortMapSchemas(m openapi3.Schemas, handler func(k string, v *openapi3.SchemaRef)) {
	// 提取 map 的键到一个切片
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// 对键进行排序
	sort.Strings(keys)

	// 通过排序后的键迭代 map，并调用处理函数
	for _, k := range keys {
		handler(k, m[k])
	}
}
