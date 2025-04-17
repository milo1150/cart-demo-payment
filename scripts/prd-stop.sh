#!/bin/sh

# Load environment
. ./load-env.sh

# Run Docker Compose with the correct environment
docker-compose -f ../deployments/prd/docker-compose.yaml down
