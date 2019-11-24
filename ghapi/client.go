package ghapi

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// AuthenticatedClient creates authenticated github client with provided token.
func AuthenticatedClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	ghClient := github.NewClient(tc)
	return ghClient
}
