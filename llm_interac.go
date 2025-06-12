package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/genai"
)

func handle_response(query string,api_key string) string {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{APIKey: api_key, Backend: genai.BackendGeminiAPI})
	if err != nil {
		log.Fatal(err)
	}
	result, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", genai.Text(query), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Text())
	return result.Text()
}
