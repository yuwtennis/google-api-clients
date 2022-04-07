package internal

import (
	"context"
	"golang.org/x/oauth2/google"
	"log"
)

// LoadDefaultCredentials load default credentials for google cloud
func LoadDefaultCredentials(ctx context.Context, scope string) *google.Credentials {
	creds, err := google.FindDefaultCredentials(ctx, scope)

	if err != nil {
		log.Fatalf("Unable to load credentials %v", err)
	}

	return creds
}
