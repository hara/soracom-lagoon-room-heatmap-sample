#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE:-$0}")" || exit
cd ..

if [ $# != 1 ]; then
    echo "Usage: ./scripts/deploy.sh YOUR_SAM_CONFIG_ENV" >&2
    exit 1
else
    echo "Deploy with --config-env $1" >&2
fi

if [ -f ./deployments/aws-sam/samconfig.toml ]; then
    sam deploy \
        --config-file ./deployments/aws-sam/samconfig.toml \
        --config-env "$1"
else
    sam deploy --guided
fi
