#!/usr/bin/env bash

# Prototype benchmark results
proto=$(cat .github/data/results.md.proto)

# Debug Message
echo "[go-orm-benchmarks] DBG: Benchmarks are running at the moment."

# Build & Run Container
docker-compose build --build-arg &>/dev/null
docker-compose up --exit-code-from benchmarker

# Apply output to template
logs=$(docker logs go-orm-benchmarks-benchmarker-1 --tail 114)
escaped=$(echo "${logs}" | sed '$!s@$@\\@g')

proto=$(sed "s|@(benchmark-results)|${escaped}|g" <<< $proto)

# Print final results
echo "$proto"