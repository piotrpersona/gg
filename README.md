# gg
![](https://github.com/piotrpersona/gg/workflows/CI/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/piotrpersona/gg)](https://goreportcard.com/report/github.com/piotrpersona/gg)
[![Documentation](https://godoc.org/github.com/piotrpersona/gg?status.svg)](http://godoc.org/github.com/piotrpersona/gg)
[![Open Source Love](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/)

Github users graph visualisation.
This project was realised during Individual Project University course.

![Contributors visualisation](svg/preview.svg?sanitize=true)

## Run

> Requiers `docker`

Start `neo4j`

> NOTE: `/plugins` and `/conf` should be mounted in order to detect clusters

```console
docker run \
    --name testneo4j \
    -p7474:7474 -p7687:7687 \
    -d \
    --env NEO4J_AUTH=neo4j/test \
    --memory 4gi \
    -v ./conf:/conf \
    -v ./plugins:/plugins \
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

## Detect communities

Enter http://localhost:7474 and pass credentials defined with `NEO_AUTH`.

Call Louvain algorithm:

```sql
CALL algo.beta.louvain.stream('User', null, { graph: 'huge', direction: 'BOTH' })
YIELD nodeId, community, communities
RETURN algo.asNode(nodeId).Name as name, community, communities
ORDER BY name ASC
```

Example result of louvain algorithm

|name              |community|communities|
|------------------|---------|-----------|
|BethGriggs        |16       |null       |
|BridgeAR          |16       |null       |
|ChALkeR           |23       |null       |
|Trott             |23       |null       |
|ZYSzys            |16       |null       |
|addaleax          |22       |null       |
|antsmartian       |22       |null       |
|bcoe              |11       |null       |
|bnoordhuis        |22       |null       |
|cjihrig           |7        |null       |
|devnexen          |16       |null       |
|devnexen          |16       |null       |
|devsnek           |22       |null       |
|devsnek           |22       |null       |
|dnlup             |16       |null       |
|eugeneo           |7        |null       |
|gabrielschulhof   |16       |null       |
|gabrielschulhof   |16       |null       |
|gireeshpunathil   |23       |null       |
|guybedford        |22       |null       |
|jasnell           |23       |null       |
|jkrems            |22       |null       |
|johnmuhl          |7        |null       |
|johnmuhl          |7        |null       |
|joyeecheung       |7        |null       |
|juanarbol         |7        |null       |
|legendecas        |7        |null       |
|ljharb            |22       |null       |
|lpinca            |16       |null       |
|lundibundi        |16       |null       |
|marswong          |23       |null       |
|mcollina          |7        |null       |
|mscdex            |7        |null       |
|mscdex            |7        |null       |
|nodejs-github-bot |16       |null       |
|nschonni          |23       |null       |
|richardlau        |16       |null       |
|ronag             |7        |null       |
|santoshyadav198613|23       |null       |
|targos            |22       |null       |
|trivikr           |23       |null       |
|wa-Nadoo          |7        |null       |
|yinzara           |7        |null       |
