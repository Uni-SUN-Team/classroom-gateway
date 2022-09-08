package list

import "unisun/api/classroom-gateway/src/models"

type ClassRoomLists struct {
	ClassRooms []ClassRoomListBody `json:"class_room"`
}

type ClassRoomListBody struct {
	ID           int                  `json:"id"`
	Thumbnail    models.Thumbnail     `json:"thumbnail"`
	Rating       ClassRoomListRating  `json:"rating"`
	Title        string               `json:"title"`
	ShortContent string               `json:"short_content"`
	Advisors     ClassRoomListAdvisor `json:"advisors"`
	Slug         string               `json:"slug"`
	Price        ClassRoomListPrice   `json:"price"`
	NumberOfUser int64                `json:"number_of_user"`
	SchoolHours  int64                `json:"school_hours"`
}

type ClassRoomListRating struct {
	Rate    int   `json:"rate"`
	Reviews int64 `json:"reviews"`
}

type ClassRoomListAdvisor struct {
	Thumbnail models.Thumbnail `json:"thumbnail"`
	FullName  string           `json:"full_name"`
	Role      string           `json:"role"`
}

type ClassRoomListPrice struct {
	RegularPrice float64 `json:"regular_price"`
	SpecialPrice float64 `json:"special_price"`
}
