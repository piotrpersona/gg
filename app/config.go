package app

import (
	"github.com/piotrpersona/gg/ghapi"
	"github.com/piotrpersona/gg/neo"
	log "github.com/sirupsen/logrus"
)

// ApplicationConfig represents application config. It includes GitHub and
// Neo4j related properties.
type ApplicationConfig struct {
	URI, Username, Password, Token string
	LogLevel                       log.Level
	PullRequestWeights             ghapi.PullRequestServicesWeights
	Query                          string
}

func (appConfig ApplicationConfig) neoconfig() neo.Config {
	return neo.Config{
		URI:      appConfig.URI,
		Username: appConfig.Username,
		Password: appConfig.Password,
	}
}
