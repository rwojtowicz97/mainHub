package internal

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Twitch struct {
	name   string
	isLive bool
	client *http.Client
}

func NewChannel(name string, client *http.Client) *Twitch {
	return &Twitch{
		name:   name,
		isLive: false,
		client: &http.Client{},
	}
}

func (t *Twitch) CheckIfLive() {
	// url := fmt.Sprintf("https://api.twitch.tv/helix/streams?user_login=" + t.name)
	url := "https://api.twitch.tv/helix/channels/followed?user_id=72164637"
	//url := "https://api.twitch.tv/helix/users?login=ejtzo"
	req, err2 := http.NewRequest("GET", url, nil)
	if err2 != nil {
		log.Printf(err2.Error())
	}

	req.Header = http.Header{
		"Authorization": {"Bearer xxxxxx"},
		"Client-Id":     {"xxxxx"},
	}

	resp, err := t.client.Do(req)
	if err != nil {
		log.Printf("http request error: %v", err)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(bodyBytes))
	fmt.Printf("\n%T\n", bodyBytes)
}
