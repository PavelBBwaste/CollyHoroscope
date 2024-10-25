package main

import (
	"fmt"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"

	"CollyHoroscope/config"
	"CollyHoroscope/horoscope"
	"CollyHoroscope/quotes"
	"CollyHoroscope/tarot"
	"CollyHoroscope/utils"
)

func main() {
	bot, err := telego.NewBot(config.BotToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println("Ошибка при создании бота:", err)
		return
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)
	bh, _ := th.NewBotHandler(bot, updates)
	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Handle(startHandler(bot), th.CommandEqual("start"))
	bh.Handle(callbackHandler(bot), th.AnyCallbackQuery())
	bh.Start()
}

func startHandler(bot *telego.Bot) th.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		keyboard := tu.InlineKeyboard(
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("♈ Овен").WithCallbackData("Овен"),
				tu.InlineKeyboardButton("♉ Телец").WithCallbackData("Телец"),
				tu.InlineKeyboardButton("♊ Близнецы").WithCallbackData("Близнецы"),
			),
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("♋ Рак").WithCallbackData("Рак"),
				tu.InlineKeyboardButton("♌ Лев").WithCallbackData("Лев"),
				tu.InlineKeyboardButton("♍ Дева").WithCallbackData("Дева"),
			),
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("♎ Весы").WithCallbackData("Весы"),
				tu.InlineKeyboardButton("♏ Скорпион").WithCallbackData("Скорпион"),
				tu.InlineKeyboardButton("♐ Стрелец").WithCallbackData("Стрелец"),
			),
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("♑ Козерог").WithCallbackData("Козерог"),
				tu.InlineKeyboardButton("♒ Водолей").WithCallbackData("Водолей"),
				tu.InlineKeyboardButton("♓ Рыбы").WithCallbackData("Рыбы"),
			),
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("🔮 Предсказание").WithCallbackData("Предсказание"),
				tu.InlineKeyboardButton("🃏 Таро").WithCallbackData("Таро"), // Добавление кнопки Таро
			),
		)

		_, _ = bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			"Выбери свой знак зодиака, нажми кнопку с предсказанием или картой Таро",
		).WithReplyMarkup(keyboard))
	}
}

func callbackHandler(bot *telego.Bot) th.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		callbackData := update.CallbackQuery.Data
		chatID := update.CallbackQuery.From.ID

		responseText, imagePath := handleCallback(callbackData)
		if imagePath != "" {
			_, err := bot.SendPhoto(&telego.SendPhotoParams{
				ChatID:  tu.ID(chatID),
				Photo:   tu.File(utils.MustOpen(imagePath)),
				Caption: responseText,
			})
			if err != nil {
				fmt.Printf("Ошибка при отправке фото: %v\n", err)
			}
		} else {
			_, _ = bot.SendMessage(tu.Message(
				tu.ID(chatID),
				responseText,
			))
		}
	}
}

func handleCallback(callbackData string) (string, string) {
	switch callbackData {
	case "Предсказание":
		return quotes.GetRandomQuote(), config.ImagePaths["Предсказание"]
	case "Таро":
		card := tarot.GetRandomCard()
		return fmt.Sprintf("%s: %s", card.Name, card.Description), card.ImagePath
	default:
		if sign, exists := config.HoroscopeSigns[callbackData]; exists {
			return horoscope.GetHoroscope(sign), config.ImagePaths["Гороскоп"]
		}
		return "Неизвестный выбор!", ""
	}
}
