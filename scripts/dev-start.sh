#!/bin/sh

# Load environment variables from the correct `.env` file path
ENV_FILE="../.env"

# Load environment variables from .env file
if [ -f "$ENV_FILE" ]; then
    # grep -v '^#' .env
    # - grep -v → Excludes lines starting with # (comments).
    # - Reads only actual environment variables from .env, ignoring comments.
    # xargs
    # - Converts the filtered output into a single-line argument format.
    # - Ensures export VAR=value works properly.
    # export $(...)
    # -  Loads all key-value pairs into environment variables.
    # export $(grep -v '^#' "$ENV_FILE" | xargs)

    # Ensure only valid lines are exported
    set -a # Automatically export variables
    . "$ENV_FILE"
    set +a # Stop automatically exporting variables
    echo "Loaded env from $ENV_FILE"
else
    echo ".env file not found at $ENV_FILE"
fi

# Run Docker Compose with the correct environment
docker-compose -f ../deployments/dev/docker-compose.yaml up
