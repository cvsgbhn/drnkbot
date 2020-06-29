package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/buger/jsonparser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	isLocal := os.Getenv("LOCAL_DRNKBOT_DEVELOPMENT")

	if len(isLocal) == 0 {
		lambda.Start(LambdaHandler)
		return
	}
	RecursiveFetchUpdatesAndRespond()
}

// RecursiveFetchUpdatesAndRespond рекурсивно cканирует новые сообщения из телеграма
// Затем оно берет полученные сообщения и отправляет на них ответ
// Сообщение может быть и пустое так как за сообщение считается toast ивент
func RecursiveFetchUpdatesAndRespond() {
	drnkbot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	telegramUpdates := tgbotapi.NewUpdate(0)
	telegramUpdates.Timeout = 60
	updates, _ := drnkbot.GetUpdatesChan(telegramUpdates)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		SendTelegramMessage(update.Message.Chat.ID)
	}
}

// LambdaHandler отвечает за обработку ивента на продакшне в AWS Lambda
// Берет json присланный телеграмом, достает оттуда chatID, отправляет сообщение в телеграм
// Возвращает статус 200 для POST от серверов телеграма, иначе они бесконечно присылают ивент
func LambdaHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	chatID, _ := jsonparser.GetInt([]byte(request.Body), "message", "chat", "id")

	SendTelegramMessage(chatID)

	return &events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

// SendTelegramMessage отправляет сообщение из функции whatthedrink в чат с id из chatID
// Так же к конструктору сообщения добавляется кнопка "Хочу попить!" из var keyboard
func SendTelegramMessage(chatID int64) tgbotapi.Message {
	drnkbot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Хочу попить!"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, createmessage(whatthedrink(), getweather()))
	msg.ReplyMarkup = keyboard
	message, err := drnkbot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
	return message
}
