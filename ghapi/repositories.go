package ghapi

import (
	"context"

	"github.com/piotrpersona/gg/model"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/neo"
)

const (
	DEFAULT_QUERY = "stars:>=1000"
)

// FetchQueriedRepositories will download repositories by the gived query.
func FetchQueriedRepositories(
	githubClient *github.Client,
	page, perPage int,
	query string) (repositories []neo.Resource, err error) {
	ctx := context.Background()
	options := github.SearchOptions{
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: perPage,
		},
	}

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
