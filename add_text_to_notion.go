package main

import (
	"context"

	notion "github.com/dstotijn/go-notion"
)



func AddTextToNotion(text string) error {
	clientNotion := notion.NewClient(notionAPIKey)

	// Specify the page ID of the Notion page where you want to add the paragraph

	// Create a new paragraph block and set its text content
	paragraph := &notion.ParagraphBlock{
		RichText: []notion.RichText{
			{
				Text: &notion.Text{
					Content: text,
				},
			},
		},
	}

	// Add the paragraph block to the specified page
	_, err := clientNotion.AppendBlockChildren(context.Background(), pageID, []notion.Block{paragraph})
	if err != nil {
		return err
	}
	return nil
}
