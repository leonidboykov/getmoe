package moebooru

import (
	"strings"
	"time"

	"github.com/leonidboykov/getmoe"
)

func (c *Client) PostsList(tags []string, page int) ([]getmoe.Post, error) {
	var posts []post
	_, err := c.sling.New().Get("posts.json").QueryStruct(queryStruct{
		Tags:  strings.Join(tags, " "),
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
			CreatedAt: time.Time(posts[i].CreatedAt),
			Author:    posts[i].Author,
			Source:    posts[i].Source,
			Rating:    posts[i].Rating,
			Hash:      posts[i].Md5,
			Tags:      posts[i].Tags,
			Score:     posts[i].Score,
		}
	}
	return result, nil
}
