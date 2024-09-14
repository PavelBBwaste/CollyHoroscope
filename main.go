package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	"github.com/gocolly/colly"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

// Структура для хранения информации о карте Таро
type TarotCard struct {
	Name        string
	Description string
	ImagePath   string
}

func main() {
	// Открытие файла с предсказаниями
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	var quotes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		quotes = append(quotes, extractQuotes(line)...)
	}

	// Получение гороскопов
	cancer := collyHoroscopeMail("cancer")
	aries := collyHoroscopeMail("aries")
	taurus := collyHoroscopeMail("taurus")
	gemini := collyHoroscopeMail("gemini")
	leo := collyHoroscopeMail("leo")
	virgo := collyHoroscopeMail("virgo")
	libra := collyHoroscopeMail("libra")
	scorpio := collyHoroscopeMail("scorpio")
	sagittarius := collyHoroscopeMail("sagittarius")
	capricorn := collyHoroscopeMail("capricorn")
	aquarius := collyHoroscopeMail("aquarius")
	pisces := collyHoroscopeMail("pisces")

	// Пути к изображениям
	imagePaths := map[string]string{
		"Гороскоп":     "images/allhorosk.JPG",
		"Предсказание": "images/zhab.JPG",
	}
	var tarotCards = []TarotCard{
		// Старшие Арканы
		{"Шут", "Начало, спонтанность, свобода, путешествие, открытие нового пути.", "images/tarot/1.png"},
		{"Маг", "Воля, мастерство, ресурсы, инициатива, создание.", "images/tarot/2.png"},
		{"Жрица", "Интуиция, тайна, подсознание, внутреннее знание, духовное развитие.", "images/tarot/3.png"},
		{"Императрица", "Плодородие, изобилие, материнство, творчество, природа.", "images/tarot/4.png"},
		{"Император", "Власть, контроль, структура, порядок, ответственность.", "images/tarot/5.png"},
		{"Жрец (Папа)", "Традиции, духовное руководство, религия, мудрость, наставничество.", "images/tarot/6.png"},
		{"Влюбленные", "Любовь, выбор, гармония, партнерство, привязанность.", "images/tarot/7.png"},
		{"Колесница", "Триумф, движение вперед, победа, сила воли, путешествие.", "images/tarot/8.png"},
		{"Сила", "Внутренняя сила, мужество, терпение, самообладание, сострадание.", "images/tarot/9.png"},
		{"Отшельник", "Одиночество, самопознание, мудрость, внутренний поиск, отстраненность.", "images/10.png"},
		{"Колесо Фортуны", "Судьба, изменение, удача, циклы, неожиданности.", "images/tarot/11.png"},
		{"Справедливость", "Справедливость, правда, ответственность, закон, честность.", "images/tarot/12.png"},
		{"Повешенный", "Жертвенность, пауза, изменение перспективы, принятие, отказ.", "images/tarot/13.png"},
		{"Смерть", "Трансформация, окончание, перерождение, изменения, конец цикла.", "images/tarot/14.png"},
		{"Умеренность", "Баланс, гармония, умеренность, терпение, адаптация.", "images/tarot/15.png"},
		{"Дьявол", "Искушение, привязанность, зависимость, ограничение, иллюзия.", "images/tarot/16.png"},
		{"Башня", "Разрушение, перемены, кризис, катастрофа, шок.", "images/tarot/17.png"},
		{"Звезда", "Надежда, вдохновение, вера, исцеление, духовное руководство.", "images/tarot/18.png"},
		{"Луна", "Иллюзия, страхи, интуиция, подсознание, загадочность.", "images/tarot/19.png"},
		{"Солнце", "Радость, успех, процветание, жизненная сила, оптимизм.", "images/tarot/20.png"},
		{"Суд", "Пробуждение, самооценка, возрождение, судьба, кармическое воздаяние.", "images/tarot/21.png"},
		{"Мир", "Завершение, достижение, гармония, завершение цикла, удовлетворение.", "images/tarot/22.png"},

		// Масть Жезлов
		{"Туз Жезлов", "Вдохновение, новые начинания, амбиции.", "images/tarot/23.png"},
		{"Двойка Жезлов", "Планирование, принятие решений, перспектива.", "images/tarot/24.png"},
		{"Тройка Жезлов", "Расширение, дальновидность, путешествие.", "images/tarot/25.png"},
		{"Четверка Жезлов", "Празднование, стабильность, достижение.", "images/tarot/26.png"},
		{"Пятерка Жезлов", "Конкуренция, борьба, конфликт.", "images/tarot/27.png"},
		{"Шестерка Жезлов", "Победа, признание, триумф.", "images/tarot/28.png"},
		{"Семерка Жезлов", "Защита, уверенность, преодоление препятствий.", "images/tarot/29.png"},
		{"Восьмерка Жезлов", "Движение, ускорение, прогресс.", "images/tarot/30.png"},
		{"Девятка Жезлов", "Настойчивость, защита, готовность к борьбе.", "images/tarot/31.png"},
		{"Десятка Жезлов", "Ответственность, бремя, давление.", "images/tarot/32.png"},
		{"Паж Жезлов", "Исследование, энтузиазм, творчество.", "images/tarot/33.png"},
		{"Рыцарь Жезлов", "Стремление, авантюризм, динамика.", "images/tarot/34.png"},
		{"Королева Жезлов", "Страсть, уверенность, харизма.", "images/tarot/35.png"},
		{"Король Жезлов", "Лидерство, предпринимательство, вдохновение.", "images/tarot/36.png"},

		// Масть Кубков
		{"Туз Кубков", "Любовь, начало новых отношений, духовное удовлетворение.", "images/tarot/37.png"},
		{"Двойка Кубков", "Партнерство, привязанность, союз.", "images/tarot/38.png"},
		{"Тройка Кубков", "Дружба, радость, празднование.", "images/tarot/39.png"},
		{"Четверка Кубков", "Апатия, разочарование, самоанализ.", "images/tarot/40.png"},
		{"Пятерка Кубков", "Потеря, сожаление, печаль.", "images/tarot/41.png"},
		{"Шестерка Кубков", "Воспоминания, ностальгия, прошлое.", "images/tarot/42.png"},
		{"Семерка Кубков", "Выбор, иллюзии, мечты.", "images/tarot/43.png"},
		{"Восьмерка Кубков", "Отход, поиск, разочарование.", "images/tarot/44.png"},
		{"Девятка Кубков", "Исполнение желаний, счастье, удовлетворение.", "images/tarot/45.png"},
		{"Десятка Кубков", "Счастье, гармония, семейное благополучие.", "images/tarot/46.png"},
		{"Паж Кубков", "Романтика, воображение, творчество.", "images/tarot/47.png"},
		{"Рыцарь Кубков", "Увлеченность, идеализм, поиски.", "images/tarot/48.png"},
		{"Королева Кубков", "Забота, интуиция, сострадание.", "images/tarot/49.png"},
		{"Король Кубков", "Эмоциональная зрелость, сочувствие, понимание.", "images/tarot/50.png"},

		// Масть Мечей
		{"Туз Мечей", "Ясность, истина, идея.", "images/tarot/51.png"},
		{"Двойка Мечей", "Выбор, сомнение, баланс.", "images/tarot/52.png"},
		{"Тройка Мечей", "Боль, печаль, разбитое сердце.", "images/tarot/53.png"},
		{"Четверка Мечей", "Отдых, восстановление, размышление.", "images/tarot/54.png"},
		{"Пятерка Мечей", "Конфликт, поражение, жертва.", "images/tarot/55.png"},
		{"Шестерка Мечей", "Перемены, переход, решение проблем.", "images/tarot/56.png"},
		{"Семерка Мечей", "Хитрость, обман, стратегия.", "images/tarot/57.png"},
		{"Восьмерка Мечей", "Ограничение, страх, беспомощность.", "images/tarot/58.png"},
		{"Девятка Мечей", "Тревога, страхи, ночные кошмары.", "images/tarot/59.png"},
		{"Десятка Мечей", "Конец, разрушение, завершение цикла.", "images/tarot/60.png"},
		{"Паж Мечей", "Наблюдение, любознательность, честность.", "images/tarot/61.png"},
		{"Рыцарь Мечей", "Стремительность, решительность, стремление к правде.", "images/tarot/62.png"},
		{"Королева Мечей", "Аналитика, ясность, справедливость.", "images/tarot/63.png"},
		{"Король Мечей", "Интеллект, власть, логика.", "images/tarot/64.png"},

		// Масть Пентаклей
		{"Туз Пентаклей", "Возможности, изобилие, процветание.", "images/tarot/65.png"},
		{"Двойка Пентаклей", "Баланс, адаптация, многозадачность.", "images/tarot/66.png"},
		{"Тройка Пентаклей", "Сотрудничество, мастерство, работа в команде.", "images/tarot/67.png"},
		{"Четверка Пентаклей", "Сдержанность, контроль, накопление.", "images/tarot/68.png"},
		{"Пятерка Пентаклей", "Потеря, трудности, нищета.", "images/tarot/69.png"},
		{"Шестерка Пентаклей", "Благотворительность, щедрость, поддержка.", "images/tarot/70.png"},
		{"Семерка Пентаклей", "Ожидание, усилия, оценка результатов.", "images/tarot/71.png"},
		{"Восьмерка Пентаклей", "Труд, усердие, совершенствование.", "images/tarot/72.png"},
		{"Девятка Пентаклей", "Самодостаточность, успех, материальное благополучие.", "images/tarot/73.png"},
		{"Десятка Пентаклей", "Богатство, наследие, семья.", "images/tarot/74.png"},
		{"Паж Пентаклей", "Ученичество, желание учиться, практичность.", "images/tarot/75.png"},
		{"Рыцарь Пентаклей", "Стабильность, преданность, трудолюбие.", "images/tarot/76.png"},
		{"Королева Пентаклей", "Заботливость, практичность, изобилие.", "images/tarot/77.png"},
		{"Король Пентаклей", "Материальное обеспечение, стабильность, лидерство.", "images/tarot/78.png"},
	}

	botToken := "..."

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println("Ошибка при создании бота:", err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)

	bh, _ := th.NewBotHandler(bot, updates)

	defer bh.Stop()
	defer bot.StopLongPolling()

	// Обработчик команды /start
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
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
	}, th.CommandEqual("start"))

	// Обработчик нажатий на кнопки
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		callbackData := update.CallbackQuery.Data
		chatID := update.CallbackQuery.From.ID

		var responseText string
		var imagePath string

		// Определение ответа и пути к изображению
		switch callbackData {
		case "Рак":
			responseText = cancer
			imagePath = imagePaths["Гороскоп"]
		case "Овен":
			responseText = aries
			imagePath = imagePaths["Гороскоп"]
		case "Телец":
			responseText = taurus
			imagePath = imagePaths["Гороскоп"]
		case "Близнецы":
			responseText = gemini
			imagePath = imagePaths["Гороскоп"]
		case "Лев":
			responseText = leo
			imagePath = imagePaths["Гороскоп"]
		case "Дева":
			responseText = virgo
			imagePath = imagePaths["Гороскоп"]
		case "Весы":
			responseText = libra
			imagePath = imagePaths["Гороскоп"]
		case "Скорпион":
			responseText = scorpio
			imagePath = imagePaths["Гороскоп"]
		case "Стрелец":
			responseText = sagittarius
			imagePath = imagePaths["Гороскоп"]
		case "Козерог":
			responseText = capricorn
			imagePath = imagePaths["Гороскоп"]
		case "Водолей":
			responseText = aquarius
			imagePath = imagePaths["Гороскоп"]
		case "Рыбы":
			responseText = pisces
			imagePath = imagePaths["Гороскоп"]
		case "Предсказание":
			responseText = quotes[rand.Intn(len(quotes))]
			imagePath = imagePaths["Предсказание"]
		case "Таро": // Обработка нажатия кнопки "Таро"
			selectedCard := tarotCards[rand.Intn(len(tarotCards))]
			responseText = fmt.Sprintf("%s: %s", selectedCard.Name, selectedCard.Description)
			imagePath = selectedCard.ImagePath
		default:
			responseText = "Неизвестный выбор!"
		}

		// Если есть изображение, отправляем его вместе с текстом
		if imagePath != "" {
			_, err := bot.SendPhoto(&telego.SendPhotoParams{
				ChatID:  tu.ID(chatID),
				Photo:   tu.File(mustOpen(imagePath)),
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
	}, th.AnyCallbackQuery())

	bh.Start()
}

// Функция для получения гороскопа
func collyHoroscopeMail(sign string) string {
	var result string
	c := colly.NewCollector(
		colly.AllowedDomains("horoscopes.rambler.ru"),
	)
	c.OnHTML("div[itemprop=articleBody]", func(h *colly.HTMLElement) {
		result = h.Text
	})
	url := fmt.Sprintf("https://horoscopes.rambler.ru/%s/", sign)
	c.Visit(url)
	return result
}

// Функция для извлечения предсказаний
func extractQuotes(text string) []string {
	var quotes []string
	for i := 0; i < len(text); i++ {
		if text[i] == '"' {
			start := i
			for j := i + 1; j < len(text); j++ {
				if text[j] == '"' && (j == len(text)-1 || text[j+1] == '.') {
					quotes = append(quotes, text[start:j+1])
					break
				}
			}
		}
	}
	return quotes
}

// Функция из документации покета, для обработки файликов картиночек
func mustOpen(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}
