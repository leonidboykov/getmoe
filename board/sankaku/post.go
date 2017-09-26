package sankaku

// Post contains native Sankaku data
type Post struct {
	Width        int         `json:"width"`
	SampleWidth  int         `json:"sample_width"`
	FileSize     int         `json:"file_size"`
	IsFavorited  bool        `json:"is_favorited"`
	Status       string      `json:"status"`
	Rating       string      `json:"rating"`
	SampleHeight int         `json:"sample_height"`
	Md5          string      `json:"md5"`
	HasComments  bool        `json:"has_comments"`
	ParentID     interface{} `json:"parent_id"`
	HasChildren  bool        `json:"has_children"`
	Change       int         `json:"change"`
	HasNotes     bool        `json:"has_notes"`
	Source       string      `json:"source"`
	Author       string      `json:"author"`
	CreatedAt    struct {
		N         int    `json:"n"`
		JSONClass string `json:"json_class"`
		S         int    `json:"s"`
	} `json:"created_at"`
	FavCount     int `json:"fav_count"`
	Height       int `json:"height"`
	PreviewWidth int `json:"preview_width"`
	Tags         []struct {
		Type   int    `json:"type"`
		NameJa string `json:"name_ja"`
		Count  int    `json:"count"`
		Name   string `json:"name"`
		ID     int    `json:"id"`
	} `json:"tags"`
	RecommendedPosts int    `json:"recommended_posts"`
	SampleURL        string `json:"sample_url"`
	FileURL          string `json:"file_url"`
	ID               int    `json:"id"`
	PreviewHeight    int    `json:"preview_height"`
	PreviewURL       string `json:"preview_url"`
	VoteCount        int    `json:"vote_count"`
	TotalScore       int    `json:"total_score"`
}

// GetWidth ...
func (p Post) GetWidth() int {
	return p.Width
}
