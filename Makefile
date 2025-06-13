
VERSION := $(shell git describe --tags --always)
COMMIT := $(shell git rev-parse --short HEAD)
DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

build:
	go build -ldflags "-X github.com/rpowelson12/Easypass/internal/version.Version=$(VERSION) -X github.com/rpowelson12/Easypass/internal/version.Commit=$(COMMIT) -X github.com/rpowelson12/Easypass/internal/version.Date=$(DATE)" -o easypass ./

install:
	go install -ldflags "-X github.com/rpowelson12/Easypass/internal/version.Version=$(VERSION) -X github.com/rpowelson12/Easypass/internal/version.Commit=$(COMMIT) -X github.com/rpowelson12/Easypass/internal/version.Date=$(DATE)" ./
