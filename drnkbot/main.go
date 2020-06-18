package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func whatthedrink() string {
	var opts  = strings.Split("чй,кф", ",")
	rand.Seed(time.Now().UnixNano())
	return opts[rand.Intn(len(opts))]
}

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Хочу попить!"),
	),
)

func main() {
	lambda.Start(Handler)
}

func Handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	log.Printf("%+v\n", request.Body)

	var result map[string]interface{}
	json.Unmarshal([]byte(request.Body), &result)

	message := result["message"].(map[string]interface{})
	chat := message["chat"].(map[string]interface{})
	chatId := int64(chat["id"].(float64))

	msg := tgbotapi.NewMessage(chatId, whatthedrink())
	msg.ReplyMarkup = numericKeyboard
	bot.Send(msg)
	return &events.APIGatewayProxyResponse{StatusCode: 200, Body: "Роскомпозор"}, nil
}
