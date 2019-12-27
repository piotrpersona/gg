#!/usr/bin/env bash

docker run \
    --name testneo4j \
    -p7474:7474 -p7687:7687 \
    -d \
    --env NEO4J_AUTH=neo4j/test \
    -v ./conf:/conf \
    -v ./plugins:/plugins \
    neo4j:latest
