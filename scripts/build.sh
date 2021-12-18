#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE:-$0}")" || exit
cd ..

sam build --template-file ./deployments/aws-sam/template.yaml
