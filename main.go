package main

import (
	"fmt"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/spf13/viper"
)

// 解析openapi的文件，生成ts的interface文件
// ```yaml
// APIFOX_PROJCET:
//   - id: 1
//     name: project1
//   - id: 2
//     name: project2
// APIFOX_TOKEN: xxx
// OUTPUT_DIR: types
// ```
// 读取apifox.yaml变量APIFOX_PROJCET列表 包含id和name
// 读取apifox.yaml变量APIFOX_TOKEN

// 迭代加载openapi文件，解析出ts的interface文件

// 项目
type Project struct {
	Id   string
	Name string
	File string
}

// 项目列表
type Projects []Project

func main() {

	// 加载yaml文件
	viper.SetConfigName("apifox")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 读取APIFOX_TOKEN
	token := viper.GetString("APIFOX_TOKEN")
	// 读取OUTPUT_DIR
	odir := viper.GetString("OUTPUT_DIR")
	otype := viper.GetString("OUTPUT_TYPE")
	// 读取APIFOX_PROJCET列表
	projects := Projects{}
	err = viper.UnmarshalKey("APIFOX_PROJCET", &projects)
	if err != nil {
		panic(err)
	}

	// 迭代加载openapi文件，解析出ts的interface文件
	// Do something with route.Operation
	for _, project := range projects {

		// 准备生成文件
		fmt.Printf("Start generate id: %s, name: %s\n", project.Id, project.Name)

		projectId := project.Id
		var openApiData []byte
		if project.File != "" {
			// 读取指定文件
			filepath := GetFilepath(project.File)
			openApiData, err = os.ReadFile(filepath)
			if err != nil {
				panic(err)
			}
		} else {
			openApiData, err = loadOpenApi(projectId, token)
			if err != nil {
				panic(err)
			}
		}

		// fmt.Printf("Load %s\n", string(openApiData))
		// 解析openapi文件
		loader := openapi3.NewLoader()
		doc, err := loader.LoadFromData(openApiData)
		if err != nil {
			panic(err)
		}
		// 生成Components name typescript文件
		filename := fmt.Sprintf("/%s.%s", project.Name, otype)

		// schemas := doc.Components.Schemas
		switch otype {
		case "ts":
			fileTypescript(odir, project.Name, doc.Components.Schemas)
		case "go":
			fileGo(odir, project.Name, doc.Components.Schemas)
		case "tsclient":
			fileTypescript(fmt.Sprintf("%s/types/", odir), fmt.Sprintf("%s.d", project.Name), doc.Components.Schemas)

			fileTsClient(fmt.Sprintf("%s/client/", odir), project.Name, doc)

			fileTsRequest(odir)

		default:
			if otype != "" {
				fmt.Printf("Unsupported output type: %s\n", otype)
				break
			}
			// fileTypescript(file, doc.Components.Schemas)
		}

		// 生成完成
		fmt.Printf("Generation completed %s\n\n", filename)
	}
}

var openaiTypesTSMaperts = map[string]string{
	"string":  "string",
	"number":  "number",
	"integer": "number",
	"boolean": "boolean",
	"array":   "Array",
	"object":  "object",
	"null":    "null",
	"any":     "any",
}

var openaiTypesGOMaperts = map[string]string{
	"string":  "string",
	"number":  "float64",
	"integer": "int64",
	"boolean": "bool",
	"array":   "[]",
	"object":  "struct",
	"null":    "interface{}",
	"any":     "interface{}",
}
