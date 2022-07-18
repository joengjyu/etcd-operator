GIT_SHA="`git rev-parse --short HEAD || echo "GitNotFound"`"
GO_LDFLAGS="-X github.com/coreos/etcd-operator/version.GitSHA=$(GIT_SHA)"
OUTPUT_DIR=`pwd`/_output/bin
IMAGE=ghcr.io/joengjyu/etcd-operator-controller

BUILD_FLAGS="-X github.com/coreos/etcd-operator/version.GitSHA=$(GIT_SHA)"

output:
	@mkdir -p "$(OUTPUT_DIR)"

test:output
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags $(BUILD_FLAGS) -o $(OUTPUT_DIR)/etcd-operator ./cmd/operator/main.go

build:output
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags $(GO_LDFLAGS) -o $(OUTPUT_DIR)/etcd-operator -installsuffix cgo ./cmd/operator/main.go
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags $(GO_LDFLAGS) -o $(OUTPUT_DIR)/etcd-backup-operator -installsuffix cgo ./cmd/backup-operator/main.go
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags $(GO_LDFLAGS) -o $(OUTPUT_DIR)/etcd-restore-operator -installsuffix cgo ./cmd/restore-operator/main.go

image: build
	@docker build -t $(IMAGE):$(GIT_SHA) -f Dockerfile $(OUTPUT_DIR)

push: image
	@docker push $(IMAGE):$(GIT_SHA)

k8s:
	chmod +x vendor/k8s.io/code-generator/generate-groups.sh
	vendor/k8s.io/code-generator/generate-groups.sh \
      "all" \
      "github.com/coreos/etcd-operator/pkg/generated" \
      "github.com/coreos/etcd-operator/pkg/apis" \
      "etcd:v1beta2" \
      --go-header-file "./hack/k8s/codegen/boilerplate.go.txt"