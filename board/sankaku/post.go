package sankaku

import (
	"time"
)

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
	RecommendedPosts interface{} `json:"recommended_posts"`
	SampleURL        string      `json:"sample_url"`
	FileURL          string      `json:"file_url"`
	ID               int         `json:"id"`
	PreviewHeight    int         `json:"preview_height"`
	PreviewURL       string      `json:"preview_url"`
	VoteCount        int         `json:"vote_count"`
	TotalScore       int         `json:"total_score"`
}

func (p *Post) parseTags() []string {
	result := make([]string, len(p.Tags))
	for i := range p.Tags {
		result[i] = p.Tags[i].Name
	}
	return result
}

func (p *Post) parseTime() time.Time {
	return time.Unix(int64(p.CreatedAt.S), 0)
}
