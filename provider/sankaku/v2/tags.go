package sankaku

type tagType int

const (
	generalType tagType = iota
	artistType
	studioType
	copyrightType
	characterType
	genreType
	_
	_
	mediumType
	metaType
)

type tag struct {
	ID     int    `json:"id"`
	NameEn string `json:"name_en"`
	NameJa string `json:"name_ja"`
	Type   int    `json:"type"`
	Count  int    `json:"count"`
	Locale string `json:"locale"`
	Rating int    `json:"rating"`
	Name   string `json:"name"`
}

func (p *post) findArtist() string {
	return filterTags(p.Tags, artistType)
}

func filterTags(tags []tag, tagType tagType) string {
	for i := range tags {
		if tags[i].Type == int(tagType) {
			return tags[i].Name
		}
	}
	return ""
}
