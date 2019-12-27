package neo

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	log "github.com/sirupsen/logrus"
)

func performQuery(session neo4j.Session, q Query) (result interface{}, err error) {
	result, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		log.Debug("Performing query: ", q)
		result, err := transaction.Run(string(q), map[string]interface{}{"": nil})
		if err != nil {
			return nil, err
		}
		log.Debug("Query result: ", result)

		if result.Next() {
			return result.Record().GetByIndex(0), nil
		}

		return nil, result.Err()
	})
	return
}

func execute(config Config, query ...Query) (err error) {
	session, driver, err := CreateSession(config)
	if err != nil {
		log.Error("Unable to create session")
		log.Error(err)
		return
	}
	defer driver.Close()
	defer session.Close()
	for _, q := range query {
		result, err := performQuery(session, q)
		if err != nil {
			log.Error("Unable to perform query: ", q)
			log.Fatal(err)
		}
		log.Debug("Query result: ", result)
	}
	return err
}
