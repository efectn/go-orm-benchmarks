#!/usr/bin/env bash

# Prototype benchmark results
proto=$(cat .github/data/results.md.proto)

# Build benchmarker
go build main.go

# Apply output to template
logs=$(./main -source="host=localhost user=postgres password=postgres dbname=test sslmode=disable" -debug=false)
escaped=$(echo "${logs}" | sed '$!s@$@\\@g')

proto=$(sed "s|@(benchmark-results)|${escaped}|g" <<< $proto)

# Print final results & delete benchmarker
rm main
echo $proto