package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	headersJson := fmt.Sprintf(`{
		"X-Apifox-Api-Version": "2024-03-28",
		"Authorization": "Bearer %s",
		"Content-Type": "application/json"
		}`, token)
	fmt.Println(req.URL.String(), "\n", headersJson, "\n", string(jsonData))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error response code: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
