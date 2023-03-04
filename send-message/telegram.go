package send_message

import (
	"awesomeProject1/secret"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)



type BotSendMessageID struct {
	Result struct{
		Message_id int
	}
}

func BotInit() *tgbotapi.BotAPI{
	bot, err := tgbotapi.NewBotAPI(secret.BOT_TOKEN)
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

func SendMessage(Team1 string, Team2 string, total_old float64,total_new float64, CID int, Id int) {
	bot := BotInit()
	//bot.Debug = true
	msg := tgbotapi.NewMessage(secret.ChatID,fmt.Sprintf("В матче %v - %v тотал изменился\n%v -> %v\nhttps://www.fon.bet/live/basketball/%v/%v",Team1,Team2,total_old,total_new,CID,Id))
	bot.Send(msg)
}


func TakeK (k float64) float64{
	bot := BotInit()
	u := tgbotapi.NewUpdate(0)
	bot.Send(tgbotapi.NewMessage(secret.ChatID,"Введите число N"))
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil { // If we got a message
			result, err := strconv.ParseFloat(update.Message.Text, 64)
			if err!=nil{
				bot.Send(tgbotapi.NewMessage(secret.ChatID,"Неверно указано число"))
			} else {
				bot.Send(tgbotapi.NewMessage(secret.ChatID,fmt.Sprintf("Число N равно %v", result)))
				k = result
				break
			}
		}
	}
	return k
}

