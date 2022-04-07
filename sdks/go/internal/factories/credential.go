package factories

import (
	"context"
	"golang.org/x/oauth2/google"
	"log"
)

type Credential struct {
	Credential *google.Credentials
}

func NewCredential() *Credential {
	c := new(Credential)

	return c
}

// Create load default credentials for google cloud
func (cred *Credential) Create(ctx context.Context, scopes []string, subjectEmail string) {

	param := new(google.CredentialsParams)
	param.Scopes = scopes
	param.Subject = subjectEmail

	creds, err := google.FindDefaultCredentialsWithParams(ctx, *param)

	if err != nil {
		log.Fatalf("Unable to load credentials %v", err)
	}

	cred.Credential = creds
}
