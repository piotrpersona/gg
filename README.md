# gg

![](https://github.com/piotrpersona/gg/workflows/CI/badge.svg?branch=master)
![](https://github.com/piotrpersona/gg/workflows/CI/badge.svg?branch=develop)

[![Go Report Card](https://goreportcard.com/badge/github.com/piotrpersona/gg)](https://goreportcard.com/report/github.com/piotrpersona/gg)
[![Documentation](https://godoc.org/github.com/piotrpersona/gg?status.svg)](http://godoc.org/github.com/piotrpersona/gg)

Github users graph visualisation.
This project was realised during Individual Project University course.

![Contributors visualisation](svg/preview.svg?sanitize=true)

## Run

> Requiers `docker`

Start `neo4j`

```console
docker run \
    --name testneo4j \
    -p7474:7474 -p7687:7687 \
    -d \
    --env NEO4J_AUTH=neo4j/test \
    --memory 4gi \
    neo4j:latest
```

or

```console
./run.sh
```

> `GITHUB_TOKEN` can be obtained at: https://github.com/settings/tokens

When `neo` is up and running run `gg`

```console
docker run \
    --rm --name gg \
    --network host \
    -e NEO_URI=bolt://localhost:7687 \
    -e NEO_USER=neo4j \
    -e NEO_PASS=test \
    -e GITHUB_TOKEN=$GITHUB_TOKEN \
    piotrpersona/gg
```

or ask for help

```console
docker run piotrpersona/gg --help
```
