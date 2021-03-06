#!/bin/bash
set -euxo pipefail

deployment_count=$1

seq 1 $deployment_count \
  | xargs -n1 -P10 -I{} \
  -- bosh -d bosh-dns-{} run-errand dns-lookuper
