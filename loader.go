package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func loadOpenApi(projectId string, token string) ([]byte, error) {

	requestBody := map[string]interface{}{"scope": map[string]interface{}{"type": "ALL"}}
	jsonData, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.apifox.com/v1/projects/%s/export-openapi", projectId), bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-Apifox-Api-Version", "2024-03-28")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")
	// Perform the request
	fmt.Println("Sending request to Apifox...")
	// headersJson := fmt.Sprintf(`{
	// 	"X-Apifox-Api-Version": "2024-03-28",
	// 	"Authorization": "Bearer %s",
	// 	"Content-Type": "application/json"
	// 	}`, token)
	// fmt.Println(req.URL.String(), "\n", headersJson, "\n", string(jsonData))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request projectId %s, error: %v\n", projectId, err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("projectId %s,	 Error response code: %d", projectId, resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 保存到文件
	// saveJson(projectId, body)

	return body, nil
}

func saveJson(projectId string, body []byte) {
	curDir, _ := os.Getwd()
	file, err := os.OpenFile(fmt.Sprintf("%s/openapi_%s.json", curDir, projectId), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(body)
	if err != nil {
		panic(err)
	}
}
