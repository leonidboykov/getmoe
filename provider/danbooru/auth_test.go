package danbooru

import (
	"testing"

	"github.com/dghubble/sling"

	"github.com/leonidboykov/getmoe"
)

func TestDanbooru_authenticate(t *testing.T) {
	tt := []struct {
		name           string
		login          string
		password       string
		hashedPassword string
		apiKey         string
		passwordSalt   string
		expectedParams string
	}{
		{
			name:  "no login",
			login: "",
		},
		{
			name:           "use api kei",
			login:          "user",
			apiKey:         "secure_api_key",
			hashedPassword: "secure_password",
			password:       "123456789",
			expectedParams: "api_key=secure_api_key&username=user",
		},
		{
			name:           "use pre-hashed password",
			login:          "user",
			hashedPassword: "secure_password",
			password:       "123456789",
			expectedParams: "login=user&password_hash=secure_password",
		},
		{
			name:           "use plain password",
			login:          "user",
			password:       "123456789",
			passwordSalt:   "choujin-steiner--%s--",
			expectedParams: "login=user&password_hash=a082648f7bcc1e40b5d562fa9b808689a36ca6be",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := danbooru{sling: sling.New()}
			d.authenticate(getmoe.Credentials{
				Login:          tc.login,
				Password:       tc.password,
				HashedPassword: tc.hashedPassword,
				APIKey:         tc.apiKey,
			}, tc.passwordSalt)
			req, err := d.sling.Request()
			if err != nil {
				t.Fatal("unexpected error", err)
			}
			if req.URL.RawQuery != tc.expectedParams {
				t.Fatal("expected", tc.expectedParams, "got", req.URL.RawQuery)
			}
		})
	}
}
