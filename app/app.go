package app

import (
	"os"
	"sync"
	"time"

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

	log.Info("Application config loaded")

	log.Info("Fetching Github repositories")
	page, err := neo.GetLastSeenPage(neoconfig)
	if err != nil {
		log.Error("Unable to fetch last seen page")
		log.Fatal(err)
		os.Exit(1)
	}

	perPage := 30
	repositories, err := ghapi.FetchQueriedRepositories(githubClient, page, perPage)
	if err != nil {
		log.Error("Unable to fetch Github repositories")
		log.Fatal(err)
		os.Exit(1)
	}

	numberOfRepositories := len(repositories)
	log.Infof("Downloaded %d repositories", numberOfRepositories)
	log.Infof("Repositories: %s", repositories)

	prRequesterService := ghapi.RequestersService{GithubClient: githubClient}
	prServices := ghapi.PullRequestServices(githubClient, appConfig.PullRequestWeights)

	var repoWg sync.WaitGroup
	repoWg.Add(numberOfRepositories)
	for _, repository := range repositories {
		repoModel := repository.(model.Repository)
		go func() {
			defer repoWg.Done()
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
					time.Sleep(time.Millisecond * 1000)
					neo.Neoize(neoconfig, prResources...)
				}
			}
		}()
	}
	repoWg.Wait()

	log.Info("gg done!")
}
