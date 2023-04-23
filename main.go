package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)



func main() {
	bot, err := tgbotapi.NewBotAPI(botApi)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil || update.Message.Voice == nil {
			continue
		}

		voice := update.Message.Voice
		err = DownloadMessage(bot, voice)
		if err != nil {
			fmt.Println("Error when downloading a file: ", err)
			return
		} else {
			fmt.Println("The file download was successful")
		}

		// Load file
		data, err := os.ReadFile("file.ogg")
		if err != nil {
			fmt.Println("Error during reading file: ", err)
			os.Remove("file.ogg")
			return
		}

		text, err := TranscriptMessage(data)
		if err != nil {
			fmt.Println("Error during transcription of the message: ", err)
			os.Remove("file.ogg")
			return
		}

		err = AddTextToNotion(text)
		if err != nil {
			log.Fatalf("Error adding paragraph: %v", err)

		} else {
			log.Println("Paragraph added successfully!")
		}
		os.Remove("file.ogg")
	}

}
