package ghapi

import (
	"context"
	"sync"

	"github.com/piotrpersona/gg/model"

	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/neo"
)

// FetchRepositories will fetch GitHub API Repositories and map them onto
// []model.Repository.
func FetchRepositories(githubClient *github.Client, limit int) (repositories []neo.Resource, err error) {
	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(limit)
	githubRepositoriesChannel := make(chan []*github.Repository, limit)
	for i := 1; i <= limit; i++ {
		since := int64(i)
		go func(wg *sync.WaitGroup, since int64, grc chan []*github.Repository) {
			defer wg.Done()
			options := &github.RepositoryListAllOptions{Since: since}
			githubRepositories, _, err := githubClient.Repositories.ListAll(ctx, options)
			if err != nil {
				return
			}
			githubRepositoriesChannel <- githubRepositories
		}(&wg, since, githubRepositoriesChannel)
	}
	wg.Wait()
	close(githubRepositoriesChannel)
	for githubRepositories := range githubRepositoriesChannel {
		for _, githubRepository := range githubRepositories {
			repository := model.CreateRepository(githubRepository)
			repositories = append(repositories, repository)
		}
	}
	return repositories, nil
}
