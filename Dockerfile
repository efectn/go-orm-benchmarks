# Define specific args for benchmark
FROM golang:alpine

# Define env vars
ARG ORM
ENV ORM=${ORM:-all}

# Build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /bin/go-orm-benchmarks ./main.go

# Run
CMD /bin/go-orm-benchmarks -source="host=postgres user=postgres password=postgres dbname=test sslmode=disable" -orm=$ORM -debug=false