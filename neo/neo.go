package neo

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func HelloWorld(uri, username, password string) (err error) {
	var (
		driver  neo4j.Driver
		session neo4j.Session
		result  neo4j.Result
	)

	driver, err = neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return
	}
	defer driver.Close()

	session, err = driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return
	}
	defer session.Close()

	_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err = transaction.Run("CREATE (p:Person)-[:LIKES]->(t:Technology)", map[string]interface{}{"": nil})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return nil, nil
		}

		return nil, result.Err()
	})
	return err
}
