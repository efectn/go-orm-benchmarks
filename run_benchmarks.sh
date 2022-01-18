#!/usr/bin/env bash

# Prototype benchmark results
proto=$(cat data/results.md.proto)

# Benchmark operations
operations=("with-no-flag|1" "multi-5|5" "multi-10|10" "multi-20|20")
#operations=("with-no-flag|")

for operation in "${operations[@]}"; do
    IFS="|"; read -a operation <<< "$operation"

    # Debug Message
    echo "[go-orm-benchmarks] DBG: Benchmarks for ${operation[0]} is running at the moment."

    # Build & Run Container
    docker-compose build --build-arg MULTI=${operation[1]} &>/dev/null
    docker-compose up --exit-code-from benchmarker

    # Apply output to template
    logs=$(docker logs go-orm-benchmarks-benchmarker-1 --tail 102) 
    escaped=$(echo "${logs}" | sed '$!s@$@\\@g')

    proto=$(sed "s|@(${operation[0]})|${escaped}|g" <<< $proto)

    echo "[go-orm-benchmarks] DBG: Benchmarks for ${operation[0]} is completed!"
done

# Print final results
echo $proto