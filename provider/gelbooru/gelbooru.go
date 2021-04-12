/*
Package gelbooru implements a simple library for accessing Gelbooru-based image
boards.
*/
package gelbooru

import (
	"github.com/dghubble/sling"
	"github.com/imdario/mergo"

	"github.com/leonidboykov/getmoe"
)

const providerName = "gelbooru"

type gelbooru struct {
	sling *sling.Sling

	postsLimit int
}

var defaultConfiguration = &getmoe.ProviderConfiguration{
	PostsLimit: 1000,
}

type apiSettings struct {
	// Force API renderer, must be `dapi`.
	page string `uri:"page"`
	// Force JSON output, must be 1.
	json int `url:"json"`
}

type queryStruct struct {
	limit int    `url:"limit"`
	tags  string `url:"tags"`
	page  int    `url:"pid"`
}

// New creates a new Gelbooru provider.
func New(config getmoe.ProviderConfiguration) getmoe.Provider {
	mergo.Merge(config, defaultConfiguration)
	g := gelbooru{
		sling: sling.New().Base(config.URL).Get("index.php").QueryStruct(apiSettings{
			page: "dapi",
			json: 1,
		}),
		postsLimit: config.PostsLimit,
	}
	g.authenticate(config.Credentials)

	return &g
}

func (g *gelbooru) RequestPage(tags getmoe.Tags, page int) ([]getmoe.Post, error) {
	var posts []post
	_, err := g.sling.New().QueryStruct(queryStruct{
		tags:  tags.String(),
		page:  page,
		limit: g.postsLimit,
	}).ReceiveSuccess(&posts)
	if err != nil {
		return nil, err
	}

	result := make([]getmoe.Post, len(posts))
	for i := range posts {
		result[i] = getmoe.Post{
			ID:        posts[i].ID,
			FileURL:   posts[i].FileURL,
			Width:     posts[i].Width,
			Height:    posts[i].Height,
			CreatedAt: posts[i].parseTime(),
			Rating:    posts[i].Rating,
			Hash:      posts[i].Hash,
			Tags:      posts[i].parseTags(),
			Score:     posts[i].Score,
		}
	}
	return result, nil
}
