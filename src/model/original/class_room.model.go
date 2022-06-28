package original

type classRoom struct {
	Id          int64              `json:"id"`
	Name        string             `json:"name"`
	CreatedAt   string             `json:"createdAt"`
	UpdatedAt   string             `json:"updatedAt"`
	PublishedAt string             `json:"publishedAt"`
	Locale      string             `json:"locale"`
	Slug        string             `json:"slug"`
	UserInClass int64              `json:"user_in_classs"`
	Price       float64            `json:"price"`
	SEO         seo                `json:"SEO"`
	Thumbnail   thumbnailLarge     `json:"thumbnail"`
	Advisors    []advisorsOriginal `json:"advisors"`
	Categories  []categories       `json:"categories"`
	Courses     []courses          `json:"Courses"`
}
