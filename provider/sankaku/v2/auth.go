package sankaku

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// authResponse holds response for token.
type authResponse struct {
	Success      bool     `json:"success"`
	TokenType    string   `json:"token_type"`
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	CurrentUser  userData `json:"current_user"`
}

// userData holds important data about current user.
type userData struct {
	ID                      int        `json:"id"`
	Name                    string     `json:"name"`
	Email                   string     `json:"email"`
	SubscriptionLevel       int        `json:"subscription_level"`
	EmailVerificationStatus string     `json:"email_verification_status"`
	IsVerified              bool       `json:"is_verified"`
	BlacklistIsHidden       bool       `json:"blacklist_is_hidden"`
	BlacklistedTags         [][]string `json:"blacklisted_tags"`
	Blacklisted             []string   `json:"blacklisted"`
	PasswordHash            string     `json:"password_hash"`
}

type credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type errorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"error"`
}

func (c *Client) authenticate(login, password string) error {
	var success authResponse
	var errorResp errorResponse

	resp, err := c.sling.New().Post("auth/token").BodyJSON(credentials{
		Login:    login,
		Password: password,
	}).Receive(&success, &errorResp)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && !errorResp.Success {
		return fmt.Errorf("sankaku: unable to authenticate: %w", errors.New(errorResp.Message))
	}

	token := strings.Join([]string{success.TokenType, success.AccessToken}, " ")
	c.sling.Set("Authorization", token)

	return nil
}
