package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.0-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(getPrompt()))
	if err != nil {
		log.Fatal(err)
	}

	printResponse(resp)
}

func getPrompt() string {
	var prompt = `
Generate a one line commit message for the following diff.
The message should be a complete sentence that simply describes the high level functional changes, with no prefixes or tags.
---
`
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		prompt += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return prompt
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
}
