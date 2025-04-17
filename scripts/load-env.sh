#!/bin/sh

# Load environment variables from the correct `.env` file path
ENV_FILE="../.env"

# Only load if the file exists
if [ -f "$ENV_FILE" ]; then
	set -a # automatically export all variables
	while IFS= read -r line || [ -n "$line" ]; do
		case "$line" in
		'' | \#*) continue ;;     # Skip comments and empty lines
		*) eval "export $line" ;; # Export key=value
		esac
	done <"$ENV_FILE"
	set +a
	echo "Loaded env from $ENV_FILE"
else
	echo ".env file not found at $ENV_FILE"
fi
