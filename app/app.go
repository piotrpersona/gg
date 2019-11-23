package app

import (
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
		contributors, err := ghapi.FetchContributors(githubClient, repoModel)
		if err != nil {
			return
		}
		pullRequests, err := ghapi.FetchPullRequests(githubClient, repoModel)
		if err != nil {
			return
		}
		issues, err := ghapi.FetchIssues(githubClient, repoModel)
		if err != nil {
			return
		}
		neo.Neoize(neoconfig, contributors...)
		neo.Neoize(neoconfig, pullRequests...)
		neo.Neoize(neoconfig, issues...)
	}
}
