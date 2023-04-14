package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Создаем новый экземпляр бота
	bot, err := tgbotapi.NewBotAPI("YOUR_BOT_TOKEN")
	if err != nil {
		log.Fatal(err)
	}

	// Обновляем seed для генерации случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Устанавливаем обработчик сообщений с типом "voice"
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
			// Генерируем случайное оскорбительное сообщение
			messages := []string{"Даун", "Соси яйца мои", "чурка тупая", "ебло тупое", "кучерявое ебло тупое", "кучерявое ебло", "уебан", "уебан тысячного ранга блять", "попуск", "попуск ебанный", "Нефор ебаный"}
			response := messages[rand.Intn(len(messages))]

			// Создаем новое сообщение в ответ на исходное сообщение
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
			msg.ReplyToMessageID = update.Message.MessageID // Устанавливаем исходное сообщение в качестве родительского
			go func() {
				_, err = bot.Send(msg)
				if err != nil {
					log.Println(err)
				}
			}()
		}
	}
}
