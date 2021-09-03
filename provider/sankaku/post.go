package sankaku

const (
	artistTag    = 1
	companyTag   = 2
	brandTag     = 3
	characterTag = 4
)

type post struct {
	ID         int      `json:"id"`
	Width      int      `json:"width"`
	Height     int      `json:"height"`
	Rating     string   `json:"rating"`
	FileSize   int      `json:"file_size"`
	FileType   string   `json:"file_type"`
	FileURL    string   `json:"file_url"`
	Source     string   `json:"source,omitempty"`
	Hash       string   `json:"md5"`
	CreatedAt  jsonTime `json:"created_at"`
	Tags       []tags   `json:"tags"`
	ParentID   int      `json:"parent_id"`
	HasNotes   bool     `json:"has_notes"`
	FavCount   int      `json:"fav_count"`
	VoteCount  int      `json:"vote_count"`
	TotalScore int      `json:"total_score"`
}

type tags struct {
	ID     int    `json:"id"`
	NameEn string `json:"name_en"`
	NameJa string `json:"name_ja"`
	Type   int    `json:"type"`
	Count  int    `json:"count"`
	Locale string `json:"locale"`
	Rating string `json:"rating"`
	Name   string `json:"name"`
}

func (p *post) parseTags() []string {
	result := make([]string, len(p.Tags))
	for i := range p.Tags {
		result[i] = p.Tags[i].Name
	}
	return result
}

func (p *post) findArtist() string {
	for i := range p.Tags {
		if p.Tags[i].Type == artistTag {
			return p.Tags[i].Name
		}
	}
	return ""
}
