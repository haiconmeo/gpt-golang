package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var gptkey string

type Request struct {
	Prompt string `json:"prompt"`
}

func gptComplete(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error parsing request body: %v", err)
		return
	}

	prompt := req.Prompt
	url := "https://api.openai.com/v1/completions"
	payload := strings.NewReader(`{
		"model": "text-davinci-003",
		"prompt": ` + "\"" + prompt + "\"" + `,
		"temperature": 1,
		"max_tokens": 256,
		"top_p": 1,
		"frequency_penalty": 0,
		"presence_penalty": 0
	  }`)

	s := fmt.Sprintf("Bearer %s", gptkey)
	reqGpt, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}
	reqGpt.Header.Set("Content-Type", "application/json")
	reqGpt.Header.Set("Authorization", s)

	client := &http.Client{}
	resp, err := client.Do(reqGpt)
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
