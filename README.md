# gg
![](https://github.com/piotrpersona/gg/workflows/CI/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/piotrpersona/gg)](https://goreportcard.com/report/github.com/piotrpersona/gg)
[![Documentation](https://godoc.org/github.com/piotrpersona/gg?status.svg)](http://godoc.org/github.com/piotrpersona/gg)
[![Open Source Love](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/)

Github users graph visualisation.
This project was realised during Individual Project University course.

![Users community](svg/preview.svg?sanitize=true)

## Design

![HLD](svg/gg-arch.svg?sanitize=true)

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
{ graph: 'huge', direction: 'BOTH', weightProperty: 'Weight', includeIntermediateCommunities: true })
YIELD nodeId, community, communities
RETURN algo.asNode(nodeId).Name as name, community, communities
ORDER BY name ASC
```

### Example result of louvain algorithm

|name              |community|communities|
|------------------|---------|-----------|
|BethGriggs        |16       |[31,16]    |
|BridgeAR          |16       |[16,16]    |
|ChALkeR           |23       |[10,23]    |
|Trott             |23       |[23,23]    |
|ZYSzys            |16       |[16,16]    |
|addaleax          |22       |[22,22]    |
|antsmartian       |22       |[22,22]    |
|bcoe              |11       |[11,11]    |
|bnoordhuis        |22       |[22,22]    |
|cjihrig           |23       |[7,23]     |
|devnexen          |16       |[31,16]    |
|devnexen          |16       |[31,16]    |
|devsnek           |22       |[15,22]    |
|devsnek           |22       |[15,22]    |
|dnlup             |16       |[16,16]    |
|eugeneo           |23       |[7,23]     |
|gabrielschulhof   |16       |[31,16]    |
|gabrielschulhof   |16       |[31,16]    |
|gireeshpunathil   |23       |[27,23]    |
|guybedford        |22       |[15,22]    |
|jasnell           |23       |[10,23]    |
|jkrems            |22       |[15,22]    |
|johnmuhl          |3        |[28,3]     |
|johnmuhl          |3        |[28,3]     |
|joyeecheung       |3        |[28,3]     |
|juanarbol         |3        |[28,3]     |
|legendecas        |23       |[7,23]     |
|ljharb            |22       |[15,22]    |
|lpinca            |16       |[16,16]    |
|lundibundi        |16       |[31,16]    |
|marswong          |23       |[23,23]    |
|mcollina          |23       |[7,23]     |
|mscdex            |3        |[28,3]     |
|mscdex            |3        |[28,3]     |
|nodejs-github-bot |16       |[31,16]    |
|nschonni          |23       |[23,23]    |
|richardlau        |16       |[31,16]    |
|ronag             |3        |[3,3]      |
|santoshyadav198613|23       |[27,23]    |
|targos            |22       |[22,22]    |
|trivikr           |23       |[27,23]    |
|wa-Nadoo          |3        |[3,3]      |
|yinzara           |23       |[7,23]     |
