package sankaku

import (
	"fmt"

	"github.com/dghubble/sling"
	"github.com/imdario/mergo"

	"github.com/leonidboykov/getmoe"
)

const providerName = "sankaku_v2"

type Client struct {
	sling *sling.Sling

	postsLimit int
}

var defaultConfiguration = &getmoe.ProviderConfiguration{
	PostsLimit: 100,
}

type queryStruct struct {
	Limit int    `url:"limit"`
	Tags  string `url:"tags"`
	Page  int    `url:"page"`
}

// type errorResp struct {
// 	Success bool     `json:"success"`
// 	Code    string   `json:"code"`
// 	Errors  []string `json:"errors"`
// }

// New creates a new Sankaku v2 provider.
func New(config getmoe.ProviderConfiguration) getmoe.Provider {
	mergo.Merge(&config, defaultConfiguration)
	c := Client{
		sling:      sling.New().Base(config.URL),
		postsLimit: config.PostsLimit,
	}
	if err := c.authenticate(config.Credentials.Login, config.Credentials.Password); err != nil {
		fmt.Println(err)
	}

	return &c
}

func (c *Client) RequestPage(tags getmoe.Tags, page int) ([]getmoe.Post, error) {
	var posts []post
	_, err := c.sling.New().Get("posts").QueryStruct(queryStruct{
		Tags:  tags.String(),
		Page:  page,
		Limit: c.postsLimit,
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
