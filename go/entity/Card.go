package entity

type Card struct {
	Id                          int    `json:"id"`
	SetName                     string `json:"setName"`
	PlayerName                  string `json:"playerName"`
	DateGraded                  string `json:"dateGraded"`
	CenteringGrade              string `json:"centeringGrade"`
	CornerGrade                 string `json:"cornerGrade"`
	EdgesGrade                  string `json:"edgesGrade"`
	SurfacesGrade               string `json:"surfacesGrade"`
	AutographGrade              string `json:"autographGrade"`
	FinalGrade                  string `json:"finalGrade"`
	TotalGradedCardsInPOPReport string `json:"totalGradedCardsInPOPReport"`
	CardsGradedAboveThisCard    string `json:"cardsGradedAboveThisCard"`
}

func (Card) TableName() string {
	return "card"

}
