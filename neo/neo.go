package neo

type Query string

type Resource interface {
	// Create Neo4j entity from a resource
	Neo() Query
}

func mapQueries(resources ...Resource) (queries []Query) {
	for _, resource := range resources {
		queries = append(queries, resource.Neo())
	}
	return
}

func Neoize(config Config, resources ...Resource) {
	queries := mapQueries(resources...)
	execute(config, queries...)
}
