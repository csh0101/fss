export GO111MODULE=on
export CGO_ENABLED=1
BUILD_TIME=$(shell date +%FT%T%z)
GIT_REVISION=$(shell git rev-parse --short HEAD)
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_BRANCH=$(shell git rev-parse --symbolic-full-name --abbrev-ref HEAD)
GIT_DIRTY=$(shell test -z "$$(git status --porcelain)" && echo "clean" || echo "dirty")
VERSION=$(shell git describe --tag --abbrev=0 --exact-match HEAD 2>/dev/null || (echo 'Git tag not found, fallback to commit id' >&2; echo ${GIT_REVISION}))

METADATA_PATH=fss/version
INJECT_VARIABLE=-X ${METADATA_PATH}.gitVersion=${VERSION} -X ${METADATA_PATH}.gitCommit=${GIT_COMMIT} -X ${METADATA_PATH}.gitBranch=${GIT_BRANCH} -X ${METADATA_PATH}.gitTreeState=${GIT_DIRTY} -X ${METADATA_PATH}.buildTime=${BUILD_TIME} -X ${METADATA_PATH}.env=${ENV}
FLAGS=-trimpath -ldflags "-s -w ${INJECT_VARIABLE}"
BINARY=fss
build-arm64:
				@echo  "> Build arm64"
				GOOS=darwin GOARCH=arm64 go build -o ./bin/${BINARY} ${FLAGS} *.go
build-amd64:
				@echo  "> Build amd64"
				GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY} ${FLAGS} *.go