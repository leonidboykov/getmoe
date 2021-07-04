package sankaku

import (
	"github.com/dghubble/sling"
	"github.com/imdario/mergo"
	"github.com/leonidboykov/getmoe"
)

const providerName = "sankaku"

type sankaku struct {
	sling *sling.Sling

	postLimit int
}

var defaultConfiguration = &getmoe.ProviderConfiguration{
	PostsLimit:   100,
	PasswordSalt: "choujin-steiner--%s--",
	AppkeySalt:   "sankakuapp_%s_Z5NE9YASej",
}

type queryStruct struct {
	limit int    `url:"limit"`
	tags  string `url:"tags"`
	page  int    `url:"page"`
}

// New creates a new Sankaku provider.
func New(config getmoe.ProviderConfiguration) getmoe.Provider {
	mergo.Merge(config, defaultConfiguration)
	s := sankaku{
		sling:     sling.New().Base(config.URL),
		postLimit: config.PostsLimit,
	}
	s.authenticate(config.Credentials, config.PasswordSalt, config.AppkeySalt)

	return &s
}

func (s *sankaku) RequestPage(tags getmoe.Tags, page int) ([]getmoe.Post, error) {
	var posts []post
	_, err := s.sling.New().Get("post/index.json").QueryStruct(queryStruct{
		tags:  tags.String(),
		page:  page,
		limit: s.postLimit,
	}).ReceiveSuccess(&posts)
	if err != nil {
		return nil, err
	}

	result := make([]getmoe.Post, len(posts))
	for i := range posts {
		result[i] = getmoe.Post{
			ID:        posts[i].ID,
			FileURL:   posts[i].FileURL,
			FileSize:  posts[i].FileSize,
			Width:     posts[i].Width,
			Height:    posts[i].Height,
			CreatedAt: posts[i].CreatedAt.Time,
			Author:    posts[i].findArtist(),
			Source:    posts[i].Source,
			Rating:    posts[i].Rating,
			Hash:      posts[i].Hash,
			Tags:      posts[i].parseTags(),
			Score:     posts[i].TotalScore,
			VoteCount: posts[i].VoteCount,
			FavCount:  posts[i].FavCount,
		}
	}
	return result, nil
}
