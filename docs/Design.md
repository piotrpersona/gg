# Github users community detection

## Scope

* Extract data from Github API
* Prepare user to user graph relation model
* Use external graph visualisation tools to analyse graph relations and detect communities

## Model

The idea is to build User nodes around the repositories they interact with.

### Github resources

Github API Link: https://developer.github.com/v3/

User `resources` that are meaningful for community relation (starting from the most important):
* Pull request review
* Pull request comment
* Issue comment
* Review request
* Issue assignees
* Followers
* Following
* Watch
* Stars
* Reactions

### Creating graph

* Given a repository fetch pull request and create node `Pull Request Author`
* For given pull request fetch related resources such as: Reviews, Pull reuqest
  comments and comments
* Connect review, pull request comment and comment authors with `Pull Request
  Author`

### Example

Reference pull request: https://github.com/microsoft/vscode/pull/16129

![PR Example](../svg/16129.svg?sanitize=true)

### Query

```sql
MERGE (requester:User {ID: 1, Name: "nojvek"})
MERGE (i:User {ID: 2, Name: "isidorn"})
MERGE (i)-[:REVIEWED {PR_ID: 16129, Weight: 20}]->(requester)
MERGE (j:User {ID: 3, Name: "jrieken"})
MERGE (j)-[:COMMENTED {PR_ID: 16129, Weight: 10}]-(requester)
RETURN requester, i, j;
```

Properties
* Node: User ID, User Name
* Relation: Pull Request ID, Weight

## Implementation

### Neo4j

Graph database which provides user interface and integrates with various graph algorithms.

Supported community detection algorithms:
* Louvain (`algo.louvain`)
* Label Propagation (`algo.labelPropagation`)
* Weakly Connected Components (`algo.unionFind`)

### Processing data

`Golang` programming language

```go
repositories := fetchRepositories()
for _, repository := range repositories {
    pullRequests := fetchPullRequests(repository)
    for _, pullRequest := range pullRequests {
        resource := fetchPRRelatedResource(pullRequest)
        neo.Create(resource)
    }
}
```

### Concurrency model

```go
repositories := fetchRepositories()

var waitGroup sync.WaitGroup
waitGroup.Add(len(repositories)) // Create task for each repository

for _, repository := range repositories {
    go func() {
        defer waitGroup.Done() // Notify waitGroup that goroutine is Done at the end
        // Processing ...
    }()
}
waitGroup.Wait() // Wait for all goroutines to be complete
```

### Docker

Container engine - used to resolve neo4j `bolt` protocol
`seabolt` dependency with Go.

### HLD

![HLD](../svg/gg-arch.svg?sanitize=true)
