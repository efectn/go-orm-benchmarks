#!/usr/bin/env bash

# Prototype benchmark results
proto=$(cat .github/data/results.md.proto)

# Build benchmarker
go build main.go

# Apply output to template
logs=$(./main -source="host=localhost user=postgres password=postgres dbname=test sslmode=disable" -debug=false)

# Print final results & delete benchmarker
rm main

echo """# Results

- orm-benchmark (with no flags)
\`\`\`
${logs}
\`\`\`""" > results.md