package factories

import (
	"context"
	"log"
	"net/http"

	"golang.org/x/oauth2/google"
)

// Client is object that holds Google sdk client object.
type Client struct {
	Client *http.Client
}

func NewClient() *Client {
  c := new(Client)

  return c
}

// Create returns Client object
func (c *Client) Create(ctx context.Context, scope string) {
	httpClient, err := google.DefaultClient(ctx, scope)

	if err == nil {
		log.Fatalf("Unable to initialize client %v", err)
	}

	c.Client = httpClient
}
