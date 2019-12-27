package app

import (
	log "github.com/sirupsen/logrus"
)

func configureLogging(level log.Level) {
	log.SetFormatter(&log.TextFormatter{
		DisableColors:    false,
		FullTimestamp:    true,
		QuoteEmptyFields: true,
	})
	log.SetLevel(level)
}
