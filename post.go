package getmoe

import "time"

// Post contains post data, represents intersection of *boorus post structs.
type Post struct {
	ID        int       `json:"id"`
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	FileURL   string    `json:"file_url"`
	FileType  string    `json:"file_type"`
	FileSize  int       `json:"file_size"`
	Tags      []string  `json:"tags"`
	Author    string    `json:"author"`
	Source    string    `json:"source"`
	Rating    string    `json:"rating"`
	Hash      string    `json:"hash"`
	Score     int       `json:"score"`
	VoteCount int       `json:"vote_count"`
	FavCount  int       `json:"fav_count"`
	CreatedAt time.Time `json:"created_at"`
}

// HasTag returns true if post has specified tag.
func (p *Post) HasTag(tag string) bool {
	for i := range p.Tags {
		if p.Tags[i] == tag {
			return true
		}
	}
	return false
}
