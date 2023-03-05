package packages

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

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
	FinalGrade                  string
	TotalGradedCardsInPOPReport string
	CardsGradedAboveThisCard    string
}

const domain = "www.beckett.com"
const baseURL = "https://www.beckett.com/grading/card-lookup?item_type=BGS&item_id="
const urlRequest = "&submit=Submit&submit=Submit"
const MIN_ITEM_ID = 15310682
const MAX_ITEM_ID = 15310782

var csvHeader = []string{
	"SetName",
	"PlayerName",
	"DateGraded",
	"CenteringGrade",
	"CornerGrade",
	"EdgesGrade",
	"SurfacesGrade",
	"AutographGrade",
	"FinalGrade",
	"TotalGradedCardsInPOPReport",
	"CardsGradedAboveThisCard",
}

// add domains
var _colly = colly.NewCollector(
	colly.AllowedDomains(domain),
)

var file *os.File

func Crawl() {
	writer := initializeCsvWriter()
	urlList := generateUrlByItemId()

	_colly.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	// look up to the html element
	_colly.OnHTML(".main_content_area > .cardDetail > tbody", func(e *colly.HTMLElement) {
		// create a map to store value
		cardMap := make(map[int]string)
		e.ForEach("tr > td:nth-child(3)", func(i int, ele *colly.HTMLElement) {
			cardMap[i] = ele.Text
		})

		data := sortMapByIndex(cardMap)
		fmt.Println("site card detail", data)

		writer.Write(data)
		// unmount the file event
		defer file.Close()
		defer writer.Flush()
	})
	for _, val := range urlList {
		time.Sleep(4 * time.Second)
		_colly.Visit(val)
	}

}

func initializeCsvWriter() *csv.Writer {
	// write into a csv file
	file, err := os.Create("beckett_card_detail.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	// header of csv file
	writer.Write(csvHeader)
	return writer
}

func sortMapByIndex(newMap map[int]string) []string {
	keys := []int{}
	row := []string{}

	for k := range newMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	if len(newMap) != 0 {
		for _, key := range keys {
			row = append(row, newMap[key])
		}
	}

	return row
}

func generateUrlByItemId() []string {
	arr := []string{}
	// may have performance issue T_T
	for i := MIN_ITEM_ID; i < MAX_ITEM_ID; i++ {
		str := []string{baseURL, strconv.Itoa(i), urlRequest}
		link := strings.Join(str, "")
		arr = append(arr, link)
	}
	return arr
}
