#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

if [ -d output ]; then
  rm -rf output
else
  echo "no output to clean"
fi
