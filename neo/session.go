package neo

import "github.com/neo4j/neo4j-go-driver/neo4j"

// CreateSession will create neo4j.Session and neo4j.Driver from provided neo4j
// config.
func CreateSession(config Config) (session neo4j.Session, driver neo4j.Driver, err error) {
	uri := config.URI
	username := config.Username
	password := config.Password
	driver, err = neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return
	}
	session, err = driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return
	}
	return
}
