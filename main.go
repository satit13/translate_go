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
	client, err := translate.NewClient(ctx, option.WithCredentialsFile("./secret/service_account.json"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	text := "my name is satit , i cannot to jumping but i am handsome "
	tag := "th"
	fmt.Printf("Text: %v\n", text)
	fmt.Printf("Translation: %v\n", makeTranslate(ctx,client,text,tag))

}
func makeTranslate(ctx context.Context,client *translate.Client,text string,tag string) string{
	// Sets the text to translate.
	// Sets the target language.
	target, err := language.Parse(tag)
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
	}

	// Translates the text into tag request.
	translations, err := client.Translate(ctx, []string{text}, target, nil)
	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
	}
	return translations[0].Text
}