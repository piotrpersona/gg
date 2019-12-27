package app

import (
	"os"
	"sync"

	"github.com/piotrpersona/gg/ghapi"
	"github.com/piotrpersona/gg/model"
	"github.com/piotrpersona/gg/neo"

	log "github.com/sirupsen/logrus"
)

// Run will run application with provided application config
func Run(appConfig ApplicationConfig) {
	configureLogging(appConfig.LogLevel)
	neoconfig := appConfig.neoconfig()
	githubClient := ghapi.AuthenticatedClient(appConfig.Token)

	repositories, err := ghapi.FetchQueriedRepositories(githubClient)
	if err != nil {
		log.Fatal("Unable to fetch Github repositories")
		log.Fatal(err)
		os.Exit(1)
	}

	prRequesterService := ghapi.RequestersService{GithubClient: githubClient}
	prServices := ghapi.PullRequestServices(githubClient, appConfig.PullRequestWeights)

	var repoWg sync.WaitGroup
	numberOfRepoTasks := len(repositories)
	repoWg.Add(numberOfRepoTasks)
	for _, repository := range repositories {
		go func() {
			defer repoWg.Done()
			repoModel := repository.(model.Repository)

			requesters, err := prRequesterService.FetchRepoResource(repoModel)
			if err != nil {
				log.Warn(err)
			}

			for _, requester := range requesters {
				requester := requester.(model.Requester)
				neo.Neoize(neoconfig, requester)
				for _, prService := range prServices {
					prResources, err := prService.Fetch(repoModel, requester.PullRequestID, requester.ID())
					if err != nil {
						log.Warn(err)
					}
					neo.Neoize(neoconfig, prResources...)
				}
			}
		}()
	}
	repoWg.Wait()
}
