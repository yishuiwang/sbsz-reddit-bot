package basic

import (
	"context"
	"fmt"
	"github.com/nstratos/go-myanimelist/mal"
	"net/http"
)

type clientIDTransport struct {
	Transport http.RoundTripper
	ClientID  string
}

func (c *clientIDTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if c.Transport == nil {
		c.Transport = http.DefaultTransport
	}
	req.Header.Add("X-MAL-CLIENT-ID", c.ClientID)
	return c.Transport.RoundTrip(req)
}

func Token() {
	publicInfoClient := &http.Client{
		// Create client ID from https://myanimelist.net/apiconfig.
		Transport: &clientIDTransport{ClientID: "xxxxxxxxxxxxxx"},
	}

	c := mal.NewClient(publicInfoClient)
	ctx := context.Background()

	// The oauth2Client will refresh the token if it expires.
	//c := mal.NewClient(oauth2Client)
	list, _, err := c.Anime.List(ctx, "one piece",
		mal.Fields{"rank", "popularity", "my_list_status"},
		mal.Limit(5),
	)
	fmt.Println(err)
	fmt.Println(list)

}
