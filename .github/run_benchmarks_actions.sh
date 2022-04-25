#!/usr/bin/env bash

# Prototype benchmark results
proto=$(cat data/results.md.proto)

# Benchmark operations
operations=("with-no-flag|1" "multi-5|5" "multi-10|10" "multi-20|20")

# Build benchmarker
go build main.go

for operation in "${operations[@]}"; do
    IFS="|"; read -a operation <<< "$operation"

    # Apply output to template
    logs=$(./main -source="host=localhost user=postgres password=postgres dbname=test sslmode=disable" -multi=${operation[1]} -debug=false) 
    escaped=$(echo "${logs}" | sed '$!s@$@\\@g')

    proto=$(sed "s|@(${operation[0]})|${escaped}|g" <<< $proto)
done

# Print final results & delete benchmarker
rm main
echo $proto