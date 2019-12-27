package neo

// Query is Neo4j Cypher compatible query string.
type Query string

// Resource represents Neo4j entity
type Resource interface {
	// Neo will create Neo4j Cypher query string from a resource.
	ID() int64
	Neo() Query
}

// Neoize will apply Resource Query on Neo4j database instance.
func Neoize(config Config, resources ...Resource) {
	queries := mapQueries(resources...)
	execute(config, queries...)
}

func mapQueries(resources ...Resource) (queries []Query) {
	for _, resource := range resources {
		queries = append(queries, resource.Neo())
	}
	return
}
