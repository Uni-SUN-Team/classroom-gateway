package card

type rating struct {
	Star       int    `json:"star"`
	Review     int64  `json:"review"`
	ReviewType string `json:"review_type"`
}
