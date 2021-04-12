package gelbooru

import (
	"fmt"

	"github.com/leonidboykov/getmoe"
)

type credentials struct {
	UserID int    `url:"user_id"`
	APIKey string `url:"api_key"`
}

func (g *gelbooru) authenticate(creds getmoe.Credentials) {
	if creds.UserID == 0 {
		fmt.Println("gelbooru: user_id is required")
		return
	}

	if creds.APIKey != "" {
		g.sling.QueryStruct(credentials{
			UserID: creds.UserID,
			APIKey: creds.APIKey,
		})
	}
}
