package neo

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

const lastSeenIDQuery = `MATCH (r:Repository) WITH max(r.ID) AS LastSeenID RETURN LastSeenID`

func FetchLastSeenID(config Config) (lastSeenID int64, err error) {
	uri := config.URI
	username := config.Username
	password := config.Password
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return
	}
	defer driver.Close()

	session, err := driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return
	}
	defer session.Close()

	result, err := performQuery(session, lastSeenIDQuery)
	if err != nil {
		return
	}
	lastSeenID = result.(int64)
	return
}
