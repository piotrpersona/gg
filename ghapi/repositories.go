package ghapi

import (
	"context"
	"fmt"

	"github.com/piotrpersona/gg/model"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/neo"
)

// FetchQueriedRepositories will download repositories by the gived query.
func FetchQueriedRepositories(githubClient *github.Client) (repositories []neo.Resource, err error) {
	ctx := context.Background()
	options := github.SearchOptions{}

	stars := 500
	// followers := 500
	// topics := 3

	query := fmt.Sprintf("stars:>=%d", stars)

	githubRepositories, _, err := githubClient.Search.Repositories(ctx, query, &options)
	if err != nil {
		return
	}
	for _, githubRepository := range githubRepositories.Repositories {
		if githubRepository.GetArchived() || githubRepository.GetPrivate() {
			continue
		}
		repository := model.CreateRepository(&githubRepository)
		repositories = append(repositories, repository)
	}
	return repositories, nil
}
