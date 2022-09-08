package purchase

import "unisun/api/classroom-gateway/src/models"

type ClassRoomPurchaseBody struct {
	ID              int                          `json:"id"`
	Slug            string                       `json:"slug"`
	Thumbnail       models.Thumbnail             `json:"thumbnail"`
	Rating          ClassRoomPurchaseRating      `json:"rating"`
	Title           string                       `json:"title"`
	ShortContent    string                       `json:"short_content"`
	Advisors        ClassRoomPurchaseAdvisor     `json:"advisors"`
	Price           ClassRoomPurchasePrice       `json:"price"`
	NumberOfUser    int64                        `json:"number_of_user"`
	SchoolHours     int64                        `json:"school_hours"`
	VideoRef        models.Video                 `json:"video_ref"`
	Content         string                       `json:"content"`
	NumberOfLessons int64                        `json:"number_of_lessons"`
	NumberOfDocs    int64                        `json:"number_of_docs"`
	Description     ClassRoomPurchaseRecommend   `json:"description"`
	Recommend       []ClassRoomPurchaseRecommend `json:"recommend"`
}

type ClassRoomPurchaseRating struct {
	Rate        int                                  `json:"rate"`
	Reviews     int64                                `json:"reviews"`
	BestOfClass []ClassRoomPurchaseRatingBestOfClass `json:"best_of_class"`
	Comments    []ClassRoomPurchaseRatingComment     `json:"comments"`
}

type ClassRoomPurchaseRatingComment struct {
	ID        int              `json:"id"`
	Thumbnail models.Thumbnail `json:"thumbnail"`
	FullName  string           `json:"full_name"`
	PostAt    string           `json:"post_at"`
	Rate      int              `json:"rate"`
	Comment   string           `json:"comment"`
}

type ClassRoomPurchaseRatingBestOfClass struct {
	Message string `json:"message"`
	Persent int    `json:"persent"`
}

type ClassRoomPurchaseAdvisor struct {
	Thumbnail models.Thumbnail `json:"thumbnail"`
	FullName  string           `json:"full_name"`
	Role      string           `json:"role"`
	Slug      string           `json:"slug"`
}

type ClassRoomPurchasePrice struct {
	RegularPrice float64 `json:"regular_price"`
	SpecialPrice float64 `json:"special_price"`
}

type ClassRoomPurchaseRecommend struct {
	Message string           `json:"message"`
	Icon    models.Thumbnail `json:"icon"`
}

type ClassRoomPurchaseLesson struct {
	Episode int    `json:"episode"`
	Name    string `json:"name"`
	Lesson  struct {
		LessonName string `json:"lesson_name"`
		Time       string `json:"time"`
	} `json:"lesson"`
}
