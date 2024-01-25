package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"openai_golang/textgenerator"
)

type RequestBody struct {
	OpenAPIKey string `json:"open_api_key"`
	Query string

}

func generateText(w http.ResponseWriter, r *http.Request, config *textgenerator.Config) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest,)
		return
	}

	params := textgenerator.OpenAIParams {
		Question: reqBody.Query,
		APIkey: reqBody.OpenAPIKey,
		MaxTokens: 1000,
		Temperature: 2.0,
	}

	responseText, err := textgenerator.GetApiResponse(params, config)

	fmt.Println(err)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error generating text: %s", err.Error()), http.StatusInternalServerError)
	return
	}
	fmt.Fprintln(w, responseText)

}
 
func main() {
	config, err := textgenerator.ReadConfig("config.json")

	if err != nil {
		fmt.Println("Error reading config file")
	}

	http.HandleFunc("/generate-text", func (w http.ResponseWriter, r *http.Request) {
		generateText(w, r, &config)
	})

	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
