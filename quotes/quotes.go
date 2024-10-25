package quotes

import (
	"bufio"
	"math/rand"
	"os"
)

var quotes []string

func init() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic("Ошибка при открытии файла с предсказаниями")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		quotes = append(quotes, scanner.Text())
	}
}

func GetRandomQuote() string {
	return quotes[rand.Intn(len(quotes))]
}
