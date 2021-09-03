package sankaku

// post is a subset of sankaku post data.
type post struct {
	ID         int      `json:"id"`
	Width      int      `json:"width"`
	Height     int      `json:"height"`
	Rating     string   `json:"rating"`
	FileURL    string   `json:"file_url"`
	FileType   string   `json:"file_type"`
	FileSize   int      `json:"file_size"`
	Source     string   `json:"source,omitempty"`
	Hash       string   `json:"md5"`
	Tags       []tag    `json:"tags"`
	ParentID   int      `json:"parent_id"`
	HasNotes   bool     `json:"has_notes"`
	FavCount   int      `json:"fav_count"`
	VoteCount  int      `json:"vote_count"`
	TotalScore int      `json:"total_score"`
	CreatedAt  jsonTime `json:"created_at"`
}

func (p *post) parseTags() []string {
	result := make([]string, len(p.Tags))
	for i := range p.Tags {
		result[i] = p.Tags[i].Name
	}
	return result
}
