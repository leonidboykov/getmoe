package danbooru

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/leonidboykov/getmoe"
)

type credentialsAPI struct {
	Login  string `url:"login"`
	APIKey string `url:"api_key"`
}

type credentials struct {
	Login        string `url:"login"`
	PasswordHash string `url:"password_hash"`
}

func (c *Client) authenticate(creds getmoe.Credentials, passwordSalt string) {
	if creds.Login == "" {
		return
	}

	if creds.APIKey != "" {
		c.sling.QueryStruct(credentialsAPI{
			Login:  creds.Login,
			APIKey: creds.APIKey,
		})
		return
	}

	if creds.HashedPassword == "" && creds.Password != "" {
		creds.HashedPassword = hashPassword(creds.Password, passwordSalt)
	}

	if creds.HashedPassword != "" {
		c.sling.QueryStruct(credentials{
			Login:        creds.Login,
			PasswordHash: creds.HashedPassword,
		})
	}
}

func hashPassword(password, salt string) string {
	password = fmt.Sprintf(salt, password)
	hash := sha1.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}
