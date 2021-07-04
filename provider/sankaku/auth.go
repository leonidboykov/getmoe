package sankaku

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/leonidboykov/getmoe"
)

type credentials struct {
	Login        string `url:"login"`
	PasswordHash string `url:"password_hash"`
	Appkey       string `url:"appkey"`
}

func (c *Client) authenticate(creds getmoe.Credentials, passwordSalt, appkeySalt string) {
	if creds.Login == "" {
		return
	}

	appkey := sha1Hash(creds.Login, appkeySalt)

	if creds.HashedPassword == "" && creds.Password != "" {
		creds.HashedPassword = sha1Hash(creds.Password, passwordSalt)
	}

	if creds.HashedPassword != "" {
		c.sling.QueryStruct(credentials{
			Login:        creds.Login,
			PasswordHash: creds.HashedPassword,
			Appkey:       appkey,
		})
	}
}

func sha1Hash(val, salt string) string {
	val = fmt.Sprintf(salt, val)
	hash := sha1.Sum([]byte(val))
	return hex.EncodeToString(hash[:])
}
