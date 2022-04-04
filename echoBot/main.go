package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5198414170:AAFPd4v3t0ty5cnSRJFN_Qjxp-e7twVk8aA")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	type M map[string]interface{}
	var UserData []M
	var Connections []M
	test := M{"User1": 1221, "User1 Text": "", "User2": 2112, "User2 Text": ""}
	Connections = append(Connections, test)
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.ID, update.Message.Text)

			newUserMsg := M{"username": update.Message.From.ID, "Text": update.Message.Text}
			// UserData = append(UserData, newUserMsg)
			if len(UserData) != 0 {
				for i := 0; i <= (len(UserData) - 1); i += 1 {
					if newUserMsg["Text"] == UserData[i]["Text"] {
						fmt.Println(UserData[i]["Text"])
						fmt.Println(newUserMsg["Text"])
						Connection := M{"User1": newUserMsg["username"], "User1 Text": newUserMsg["Text"], "User2": UserData[i]["username"], "User2 Text": UserData[i]["Text"]}

						Connections = append(Connections, Connection)
						fmt.Println(Connections)

					}
				}
				if len(Connections) == 0 {
					UserData = append(UserData, newUserMsg)
				} else if newUserMsg["username"] != Connections[len(Connections)-1]["User1"] {
					UserData = append(UserData, newUserMsg)
				}

				// for i := 0; i <= (len(UserData) - 1); i += 1 {
				// 	fmt.Println(i)
				// }
				// UserData = append(UserData, newUserMsg)

			} else {
				UserData = append(UserData, newUserMsg)
			}
			if len(Connections) > 0 {
				var User1 int = int64(Connections[0]["User1"].(int))
				var User2 int = int64(Connections[0]["User2"].(int))
				if update.Message.Chat.ID == User1{

					// a := []interface{}{1, 2, 3, 4, 5}
					// b := make([]int, len(a))
					// for i := range a {
					// 		b[i] = a[i].(int)
					// }
					// fmt.Println(a, b)

					msg := tgbotapi.NewMessage(User2, Connections[0]["User1 Text"].(string))
					bot.Send(msg)
					// msg.ReplyToMessageID = update.Message.MessageID
				} else if update.Message.Chat.ID == User2 {
					msg := tgbotapi.NewMessage(User1, Connections[0]["User2 Text"].(string))
					bot.Send(msg)
					// msg.ReplyToMessageID = update.Message.MessageID
				}else{
					fmt.Println("No Connected Users Yet")
				}
			}

		}
	}
}
