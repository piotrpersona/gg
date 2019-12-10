# github-community-visualisation

GitHub users visualisation on graph
This document is meant to be presented with [mark.show](https://mark.show/#/).

---

### Scope

* Extract data from Github API
* Build graph of users relations (e.g.: connect two users if they contribute to the same project)
* Use external graph visualisation tools to analyse graph relations and detect communities

---

### Extracting data

* There are various resources that can represent users relations, e.g.: project contribution, followed/followers, stars, issues reported, etc.
* Data extractor should provide a way to choose relation details - as it was mentioned above, the tool should build graph using some “basic” resources analysis (such as contributions, followers, followed) and allow to extend graph model including additional resources (stars, issues and more).

---

### GitHub API

* Pagination: by default 30 records per page
* Github exposes navigating through pages
    * Requests limit per hour:
    * 5000 Basic Authentication / OAuth
* 60 for unauthenticated
* Requests should be made in parallel to speed up extracting process

---

### Graph

* Neo4j - single node in memory, extensible with community detection algorithms (Louvain, Label Propagation, Weakly Connected Components)
* Dgraph - distributed, lack of algorithm plugins

---

### High Level Design

![High Level Design diagram](../svg/hld.svg?sanitize=true)

---

### Stack

* Golang for extracting data
* Python for analysis
