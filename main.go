package main

import (
	_ "fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"math/rand"
	"time"
)

func main() {
	// Create a new bot instance
	bot, err := tgbotapi.NewBotAPI(Token)
	if err != nil {
		log.Fatal(err)
	}

	// Update the random seed
	rand.Seed(time.Now().UnixNano())

	// Set up a message handler for voice messages
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Voice != nil {
			// Generate a random offensive message
			messages := []string{"Даун", "Соси яйца мои", "чурка тупая", "ебло тупое", "кучерявое ебло тупое", "кучерявое ебло", "уебан", "уебан тысячного ранга блять", "попуск", "попуск ебанный", "Нефор ебаный"}
			response := messages[rand.Intn(len(messages))]

			// Reply with the generated message
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
			msg.ReplyToMessageID = update.Message.MessageID
			_, err = bot.Send(msg)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
