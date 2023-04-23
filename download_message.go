package main

import (
	"io"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func DownloadMessage(bot *tgbotapi.BotAPI, voice *tgbotapi.Voice) error {

	fileID := voice.FileID

	fileConfig := tgbotapi.FileConfig{
		FileID: fileID,
	}

	file, err := bot.GetFile(fileConfig)
	if err != nil {
		log.Panic(err)
	}

	fileURL := file.Link(bot.Token)

	// Открытие HTTP-соединения с URL-адресом
	resp, err := http.Get(fileURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Создание локального файла
	filelocal, err := os.Create("file.ogg")
	if err != nil {
		return err
	}
	defer filelocal.Close()

	// Копирование данных из HTTP-ответа в локальный файл
	_, err = io.Copy(filelocal, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
