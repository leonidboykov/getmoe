package sankaku

import (
	"github.com/dghubble/sling"
	"github.com/imdario/mergo"
	"github.com/leonidboykov/getmoe"
)

const providerName = "sankaku"

type Client struct {
	sling *sling.Sling

	postLimit int
}

var defaultConfiguration = &getmoe.ProviderConfiguration{
	PostsLimit:   100,
	PasswordSalt: "choujin-steiner--%s--",
}

type queryStruct struct {
	limit int    `url:"limit"`
	tags  string `url:"tags"`
	page  int    `url:"page"`
}

// New creates a new Sankaku provider.
func New(config getmoe.ProviderConfiguration) getmoe.Provider {
	mergo.Merge(config, defaultConfiguration)
	c := Client{
		sling:     sling.New().Base(config.URL),
		postLimit: config.PostsLimit,
	}
	c.authenticate(config.Credentials, config.PasswordSalt)

	return &c
}

func (c *Client) RequestPage(tags getmoe.Tags, page int) ([]getmoe.Post, error) {
	var posts []post
	_, err := c.sling.New().Get("post/index.json").QueryStruct(queryStruct{
		tags:  tags.String(),
		page:  page,
		limit: c.postLimit,
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
