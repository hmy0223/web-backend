package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type CardDetail struct {
	SetName                     string
	PlayerName                  string
	DateGraded                  string
	CenteringGrade              string
	CornerGrade                 string
	EdgesGrade                  string
	SurfacesGrade               string
	AutographGrade              string
	FinalGrade                  float32
	TotalGradedCardsInPOPReport int
	CardsGradedAboveThisCard    int
}

func main() {
	crawl()
}

func crawl() {
	// write into a csv file
	file, err := os.Create("beckett_card_detail.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	// header of csv file
	headers := []string{"SetName"}
	writer.Write(headers)

	// add domains
	c := colly.NewCollector(
		colly.AllowedDomains("www.beckett.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	// look up to html element
	c.OnHTML(".cardDetail tbody > tr:nth-child(1) > td:nth-child(3)", func(e *colly.HTMLElement) {
		detail := CardDetail{}
		detail.SetName = e.Text
		fmt.Println(e.Text)
		row := []string{detail.SetName}
		writer.Write(row)
	})

	// dynamic serial number will be added later
	c.Visit("https://www.beckett.com/grading/card-lookup?item_type=BGS&item_id=15310682&submit=Submit&submit=Submit")
}
