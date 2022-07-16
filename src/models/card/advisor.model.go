package card

type advisor struct {
	Name      string         `json:"name"`
	Type      string         `json:"type"`
	Thumbnail thumbnailSmall `json:"thumbnail"`
}
