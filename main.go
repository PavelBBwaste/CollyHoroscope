package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
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

	fmt.Println(cancer)
	fmt.Println(aries)
	fmt.Println(taurus)
	fmt.Println(gemini)
	fmt.Println(leo)
	fmt.Println(virgo)
	fmt.Println(libra)
	fmt.Println(scorpio)
	fmt.Println(sagittarius)
	fmt.Println(capricorn)
	fmt.Println(aquarius)
	fmt.Println(pisces)
}

func collyHoroscopeMail(sign string) string {
	var result string
	c := colly.NewCollector(
		colly.AllowedDomains("horo.mail.ru"),
	)
	c.OnHTML(".article__text", func(h *colly.HTMLElement) {
		result = h.Text
	})
	url := fmt.Sprintf("https://horo.mail.ru/prediction/%s/today/", sign)
	c.Visit(url)
	return result
}
