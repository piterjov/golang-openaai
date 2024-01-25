package textgenerator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)
type Message struct {
	Content string `json:"content"`
}
// Choice represents a single choice in the OpenAI API response
type Choice struct {
	Message Message `json:"message"`

}

type Config struct {
	ApiEndpoint string `json:apiEndpoint`
	ApiKey string `json:apiKey`
}

type ApiResponse struct {
	Choices []Choice `json:"choices"`
}

func ReadConfig(filePath string) (Config, error) {
	var config Config
	configFile, err := os.Open(filePath)

	if err != nil {
		return config, err
	}

	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	if err != nil {
		return config, err
	}

	return config, nil
}

type OpenAIParams struct {
	Question string 
	APIkey string
	MaxTokens int
	Temperature float64

}


func GetApiResponse(params OpenAIParams, config *Config) (string, error) {
 	apiUrl := config.ApiEndpoint

 
	
 
    payload := map[string]interface{}{
        // "model": "gpt-4-1106-preview", // Replace with the correct model name
		"model": "gpt-3.5-turbo-1106", // Replace with the correct model name
        "messages": []map[string]string{
            {
                "role": "user",
                "content": params.Question,
            },
        },
    }

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	// apiKey :=  config.ApiKey

	req.Header.Set("Authorization", "Bearer "+ params.APIkey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-OK status code %d", resp.StatusCode) // TODO: better error message here
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println("API response: ", string(body))

	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	
	if err != nil {
		return "", err
	}

	if len(apiResponse.Choices) > 0 {
		return apiResponse.Choices[0].Message.Content, nil
	}
	return "", fmt.Errorf("no choices in response")
}
