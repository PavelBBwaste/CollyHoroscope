package horoscope

import (
	"fmt"

	"github.com/gocolly/colly"
)

func GetHoroscope(sign string) string {
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
