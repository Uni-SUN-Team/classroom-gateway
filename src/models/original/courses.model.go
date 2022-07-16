package original

type courses struct {
	Id          int64  `json:"id"`
	NameSubject string `json:"name_subject"`
	Courses     []data `json:"courses"`
}

type data struct {
	Id      int64  `json:"id"`
	Preview bool   `json:"preview"`
	Course  course `json:"course"`
}

type course struct {
	Id              int64              `json:"id"`
	Title           string             `json:"title"`
	Description     string             `json:"description"`
	Content         string             `json:"content"`
	Short_content   string             `json:"short_content"`
	CreatedAt       string             `json:"createdAt"`
	UpdatedAt       string             `json:"updatedAt"`
	PublishedAt     string             `json:"publishedAt"`
	Locale          string             `json:"locale"`
	Slug            string             `json:"slug"`
	Thumbnail       thumbnailLarge     `json:"thumbnail"`
	CourseHighLight video              `json:"course_high_light"`
	Advisors        []advisorsOriginal `json:"advisors"`
	Categories      []categories       `json:"categories"`
	SEO             seo                `json:"SEO"`
	VideoContent    video              `json:"video_content"`
	VideoPreview    video              `json:"video_preview"`
	Attachment      []thumbnailLarge   `json:"file"`
}
