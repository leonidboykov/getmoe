package gelbooru

import (
	"github.com/leonidboykov/getmoe"
)

type credentials struct {
	UserID int    `url:"user_id"`
	APIKey string `url:"api_key"`
}

func (c *Client) authenticate(creds getmoe.Credentials) {
	if creds.UserID == 0 {
		return
	}

	if creds.APIKey != "" {
		c.sling.QueryStruct(credentials{
			UserID: creds.UserID,
			APIKey: creds.APIKey,
		})
	}
}
