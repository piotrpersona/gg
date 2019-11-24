package app

import (
	"os"

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
	repositories, err := ghapi.FetchRepositories(githubClient, appConfig.Limit)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	neo.Neoize(neoconfig, repositories...)

	repoResources := []ghapi.RepoResource{
		ghapi.Contributors{},
		ghapi.PullRequests{},
		ghapi.Issues{},
	}

	for _, repository := range repositories {
		repoModel := repository.(model.Repository)
		for _, repoResource := range repoResources {
			resources, err := repoResource.Fetch(githubClient, repoModel)
			if err != nil {
				log.Error(err)
			}
			neo.Neoize(neoconfig, resources...)
		}
	}
}
