# !/bin/bash

# Define the directories to exclude from testing
EXCLUDE_DIRS=("postgres/boilentity" "docs" "cmd" "postgres")

# Build the list of packages to test, excluding the specified directories
PACKAGES=$(go list ./... | grep -vE "$(IFS=\|; echo "${EXCLUDE_DIRS[*]}")")

go test -v $PACKAGES -count=1