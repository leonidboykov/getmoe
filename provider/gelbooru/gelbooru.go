package gelbooru

import (
	"time"

	"github.com/dghubble/sling"
	"github.com/imdario/mergo"

	"github.com/leonidboykov/getmoe"
)

const providerName = "gelbooru"

type Client struct {
	sling *sling.Sling

	postsLimit int
}

var defaultConfiguration = &getmoe.ProviderConfiguration{
	PostsLimit: 1000,
}

// apiSettings contains params to enable JSON API.
var apiSettings = struct {
	Page string `url:"page"`
	JSON int    `url:"json"`
	S    string `url:"s"`
	Q    string `url:"q"`
}{
	Page: "dapi",
	JSON: 1,
	S:    "post",
	Q:    "index",
}

type queryStruct struct {
	Limit int    `url:"limit"`
	Tags  string `url:"tags"`
	Page  int    `url:"pid"`
}

// New creates a new Gelbooru provider.
func New(config getmoe.ProviderConfiguration) getmoe.Provider {
	mergo.Merge(&config, defaultConfiguration)
	c := Client{
		sling:      sling.New().Base(config.URL).Get("index.php").QueryStruct(apiSettings),
		postsLimit: config.PostsLimit,
	}
	c.authenticate(config.Credentials)
	return &c
}

func (c *Client) RequestPage(tags getmoe.Tags, page int) ([]getmoe.Post, error) {
	var posts []post
	_, err := c.sling.New().QueryStruct(queryStruct{
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
			Width:     posts[i].Width,
			Height:    posts[i].Height,
			CreatedAt: time.Time(posts[i].Change),
			Rating:    posts[i].Rating,
			Hash:      posts[i].Hash,
			Tags:      posts[i].Tags,
			Score:     posts[i].Score,
		}
	}
	return result, nil
}
