/*
Package danbooru implements a simple library for accessing Danbooru-based image
boards.

Source code of Danbooru is available at https://github.com/r888888888/danbooru
*/
package danbooru

import (
	"github.com/dghubble/sling"
	"github.com/imdario/mergo"

	"github.com/leonidboykov/getmoe"
)

const providerName = "danbooru"

type Client struct {
	sling *sling.Sling

	passwordSalt string
	postsLimit   int
}

var defaultConfiguration = &getmoe.ProviderConfiguration{
	PasswordSalt: "choujin-steiner--%s--",
	PostsLimit:   200,
}

type queryStruct struct {
	Limit int    `url:"limit"`
	Tags  string `url:"tags"`
	Page  int    `url:"page"`
}

// New creates a new Danbooru provider.
func New(config getmoe.ProviderConfiguration) getmoe.Provider {
	mergo.Merge(config, defaultConfiguration)
	c := Client{
		sling:        sling.New().Base(config.URL),
		passwordSalt: config.PasswordSalt,
		postsLimit:   config.PostsLimit,
	}
	c.authenticate(config.Credentials, config.PasswordSalt)

	return &c
}

func (c *Client) RequestPage(tags getmoe.Tags, page int) ([]getmoe.Post, error) {
	var posts []post
	_, err := c.sling.New().Get("posts.json").QueryStruct(queryStruct{
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
			Width:     posts[i].ImageWidth,
			Height:    posts[i].ImageHeight,
			CreatedAt: posts[i].CreatedAt,
			Author:    posts[i].TagStringArtist,
			Source:    posts[i].Source,
			Rating:    posts[i].Rating,
			Hash:      posts[i].Md5,
			Tags:      posts[i].parseTags(),
			Score:     posts[i].Score,
		}
	}
	return result, nil
}
