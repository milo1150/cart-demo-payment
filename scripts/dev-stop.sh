#!/bin/sh

# Load environment
. ./load-env.sh

# Run Docker Compose with the correct environment
docker-compose -f ../deployments/dev/docker-compose.yaml down
