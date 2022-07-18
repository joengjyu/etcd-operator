#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

chmod +x vendor/k8s.io/code-generator/generate-groups.sh
vendor/k8s.io/code-generator/generate-groups.sh \
  "all" \
  "github.com/coreos/etcd-operator/pkg/generated" \
  "github.com/coreos/etcd-operator/pkg/apis" \
  "etcd:v1beta2" \
  --go-header-file "./hack/k8s/codegen/boilerplate.go.txt" \
  $@
