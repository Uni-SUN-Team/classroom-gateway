package card

type data struct {
	Id                int64   `json:"id"`
	Name              string  `json:"name"`
	ShortContent      string  `json:"short_content"`
	StudentNumber     int64   `json:"student_number"`
	StudentNumberType string  `json:"student_number_type"`
	Time              float64 `json:"time"`
	TimeType          string  `json:"time_type"`
	Advisor           advisor `json:"advisor"`
	Price             price   `json:"price"`
	Rating            rating  `json:"rating"`
}
