#!/usr/bin/env bash

# Prototype benchmark results
proto=$(cat data/results.md.proto)

# Benchmark operations
operations=("with-no-flag|1" "multi-5|5" "multi-10|10" "multi-20|20")
#operations=("with-no-flag|1")

# Build benchmarker
go build main.go

for operation in "${operations[@]}"; do
    IFS="|"; read -a operation <<< "$operation"

    # Debug Message
    echo "[go-orm-benchmarks] DBG: Benchmarks for ${operation[0]} is running at the moment."

    # Apply output to template
    logs=$(./main -source="host=localhost user=postgres password=postgres dbname=test sslmode=disable" -multi=${operation[1]} -debug=false) 
    escaped=$(echo "${logs}" | sed '$!s@$@\\@g')

    proto=$(sed "s|@(${operation[0]})|${escaped}|g" <<< $proto)

    echo "[go-orm-benchmarks] DBG: Benchmarks for ${operation[0]} is completed!"
done

# Print final results & delete benchmarker
rm main
echo $proto