package app

import (
	"github.com/google/go-github/github"
	"github.com/piotrpersona/gg/ghapi"
	"github.com/piotrpersona/gg/model"
	"github.com/piotrpersona/gg/neo"

	log "github.com/sirupsen/logrus"
)

type ApplicationConfig struct {
	URI, Username, Password, Token string
	Limit                          int
	LogLevel                       log.Level
}

func configureLogging(level log.Level) {
	log.SetFormatter(&log.TextFormatter{
		DisableColors:    true,
		FullTimestamp:    true,
		QuoteEmptyFields: true,
	})
	log.SetLevel(level)
}

type connector = func(*github.Client, model.Repository) ([]neo.Resource, error)

func connect(neoconfig neo.Config, githubClient *github.Client, c connector, repoModel model.Repository) error {
	resources, err := c(githubClient, repoModel)
	if err != nil {
		return err
	}
	neo.Neoize(neoconfig, resources...)
	return nil
}

func Run(appConfig ApplicationConfig) {
	configureLogging(appConfig.LogLevel)
	neoconfig := neo.Config{
		URI:      appConfig.URI,
		Username: appConfig.Username,
		Password: appConfig.Password,
	}
	githubClient := ghapi.AuthenticatedClient(appConfig.Token)
	repositories, err := ghapi.FetchRepositories(githubClient, appConfig.Limit)
	if err != nil {
		panic(err)
	}
	neo.Neoize(neoconfig, repositories...)
	for _, repo := range repositories {
		repoModel := repo.(model.Repository)
		connectors := []connector{
			ghapi.FetchContributors,
			ghapi.FetchPullRequests,
			ghapi.FetchIssues,
		}
		for _, connector := range connectors {
			connect(neoconfig, githubClient, connector, repoModel)
		}
	}
}
