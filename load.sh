#!/usr/bin/env bash

while true; do
    docker run \
        --network host \
        --rm --name gg \
        -e NEO_URI=bolt://localhost:7687 \
        -e NEO_USER=neo4j \
        -e NEO_PASS=test \
        -e GITHUB_TOKEN=$GITHUB_TOKEN \
        gg --loglevel info | tee /tmp/gg-fetch.log
    sleep 5m
done
