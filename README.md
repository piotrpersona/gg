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
CALL algo.beta.louvain.stream('User', null,
{ graph: 'huge', direction: 'BOTH', weightProperty: 'Weight' })
YIELD nodeId, community, communities
RETURN algo.asNode(nodeId).Name as name, community, communities
ORDER BY name ASC
```

Example result of louvain algorithm

|name              |community|
|------------------|---------|
|BethGriggs        |16       |
|BridgeAR          |16       |
|ChALkeR           |23       |
|Trott             |23       |
|ZYSzys            |16       |
|addaleax          |22       |
|antsmartian       |22       |
|bcoe              |11       |
|bnoordhuis        |22       |
|cjihrig           |7        |
|devnexen          |16       |
|devnexen          |16       |
|devsnek           |22       |
|devsnek           |22       |
|dnlup             |16       |
|eugeneo           |7        |
|gabrielschulhof   |16       |
|gabrielschulhof   |16       |
|gireeshpunathil   |23       |
|guybedford        |22       |
|jasnell           |23       |
|jkrems            |22       |
|johnmuhl          |7        |
|johnmuhl          |7        |
|joyeecheung       |7        |
|juanarbol         |7        |
|legendecas        |7        |
|ljharb            |22       |
|lpinca            |16       |
|lundibundi        |16       |
|marswong          |23       |
|mcollina          |7        |
|mscdex            |7        |
|mscdex            |7        |
|nodejs-github-bot |16       |
|nschonni          |23       |
|richardlau        |16       |
|ronag             |7        |
|santoshyadav198613|23       |
|targos            |22       |
|trivikr           |23       |
|wa-Nadoo          |7        |
|yinzara           |7        |
