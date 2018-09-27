#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

vendor/k8s.io/code-generator/generate-groups.sh \
deepcopy \
github.com/cezarsa/chaos/pkg/generated \
github.com/cezarsa/chaos/pkg/apis \
chaos:v1beta1 \
--go-header-file "./tmp/codegen/boilerplate.go.txt"
