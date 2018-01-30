// Sample translate-quickstart translates "Hello, world!" into Russian.
package main

import (
	"fmt"
	"log"

	// Imports the Google Cloud Translate client package.
	"cloud.google.com/go/translate"
	"golang.org/x/net/context"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	// Creates a client.
	client, err := translate.NewClient(ctx,option.WithCredentialsFile("./secret/service_account.json"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the text to translate.
	text := "Hello, world!"
	// Sets the target language.
	target, err := language.Parse("zh-CN")
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
	}

	// Translates the text into Russian.
	translations, err := client.Translate(ctx, []string{text}, target, nil)
	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
	}

	fmt.Printf("Text: %v\n", text)
	fmt.Printf("Translation: %v\n", translations[0].Text)
}