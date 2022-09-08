package purchased

import "unisun/api/classroom-gateway/src/models"

type ClassRoomPurchasedBody struct {
	ID             int                        `json:"id"`
	CourseIdActive int                        `json:"course_id_active"`
	Advisor        ClassRoomPurchasedAdvisor  `json:"advisor"`
	ClassVideo     models.Video               `json:"class_video"`
	Proguess       int                        `json:"proguess"`
	Courses        []ClassRoomPurchasedCourse `json:"courses"`
	Attachment     []models.Thumbnail         `json:"attachment"`
	TimePlay       string                     `json:"time_play"`
}

type ClassRoomPurchasedAdvisor struct {
	FullName string `json:"full_name"`
}

type ClassRoomPurchasedCourse struct {
	Title   string                            `json:"title"`
	Courses []ClassRoomPurchasedCourseEpisode `json:"courses"`
}

type ClassRoomPurchasedCourseEpisode struct {
	ID      int    `json:"id"`
	Episode int    `json:"episode"`
	Title   string `json:"title"`
	Lesson  string `json:"lesson"`
	Slug    string `json:"slug"`
	Active  bool   `json:"active"`
}
