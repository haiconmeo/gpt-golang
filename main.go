package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var gptkey string

func gptComplete(w http.ResponseWriter, r *http.Request) {
	url := "https://api.openai.com/v1/completions"

	data := map[string]interface{}{
		"model":       "text-davinci-003",
		"prompt":      "Say this is a test",
		"max_tokens":  7,
		"temperature": 0,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer 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")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("lỗi r")
	}

	w.Write(body)
}
func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	gptkey = os.Getenv("GPT_KEY")
	server := http.NewServeMux()
	server.HandleFunc("/gpt", gptComplete)
	err = http.ListenAndServe(":80", server)
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}
