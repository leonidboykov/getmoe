package gelbooru

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/leonidboykov/getmoe"
)

func TestGelbooru_authenticate(t *testing.T) {
	tt := []struct {
		name           string
		userID         int
		apiKey         string
		expectedParams string
	}{
		{
			name:   "no user id",
			userID: 0,
		},
		{
			name:   "no api key",
			userID: 123456,
			apiKey: "",
		},
		{
			name:           "success",
			userID:         123456,
			apiKey:         "user_api_access_key",
			expectedParams: "api_key=user_api_access_key&user_id=123456",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			g := gelbooru{sling: sling.New()}
			g.authenticate(getmoe.Credentials{
				UserID: tc.userID,
				APIKey: tc.apiKey,
			})
			req, err := g.sling.Request()
			if err != nil {
				t.Fatal("unexpected error", err)
			}
			if req.URL.RawQuery != tc.expectedParams {
				t.Fatal("expected", tc.expectedParams, "got", req.URL.RawQuery)
			}
		})
	}
}
