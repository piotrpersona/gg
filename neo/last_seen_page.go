package neo

// IncrementPageNumber will increment last seen github search api page number.
func IncrementPageNumber(config Config) (err error) {
	session, driver, err := CreateSession(config)
	if err != nil {
		return
	}
	defer driver.Close()
	defer session.Close()
	query := Query(`MATCH (p:Page) SET p.Number = p.Number + 1`)
	_, err = performQuery(session, query)
	return
}

// GetLastSeenPage will resolve last seen github search api page number.
func GetLastSeenPage(config Config) (lastSeenID int, err error) {
	session, driver, err := CreateSession(config)
	if err != nil {
		return
	}
	defer driver.Close()
	defer session.Close()

	query := Query(`MATCH (p:Page) RETURN p.Number`)
	result, err := performQuery(session, query)
	if err != nil {
		return
	}
	if result == nil {
		lastSeenID = 0
	} else {
		lastSeenID = result.(int)
	}
	return
}
