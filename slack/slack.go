package slack

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Slack struct {
	Client    *http.Client
	AuthToken string
}

func New(token string) *Slack {
	return &Slack{
		Client:    http.DefaultClient,
		AuthToken: token,
	}
}

type Profile struct {
	StatusEmoji      string `json:"status_emoji"`
	StatusText       string `json:"status_text"`
	StatusExpiration int64  `json:"status_expiration"`
}

type Response struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

func (s *Slack) SetStatus(emoji, text string, expires int) error {
	expiresUnix := time.Now().Add(time.Duration(expires) * time.Minute).Unix()

	p := Profile{StatusEmoji: emoji, StatusText: text, StatusExpiration: expiresUnix}
	pJson, err := json.Marshal(p)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("https://slack.com/api/users.profile.set?profile=%s", url.QueryEscape(string(pJson))), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json ; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.AuthToken))

	res, err := s.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var jsonRes Response
	err = json.Unmarshal(body, &jsonRes)
	if err != nil {
		return err
	}

	if !jsonRes.Ok {
		return fmt.Errorf("slack error: %s", jsonRes.Error)
	}

	return nil
}
