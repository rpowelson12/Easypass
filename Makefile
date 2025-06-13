VERSION := $(shell git describe --tags --always)

build:
	go build -ldflags "-X github.com/rpowelson12/Easypass/internal/version.Version=$(VERSION)"

install:
	go install -ldflags "-X github.com/rpowelson12/Easypass/internal/version.Version=$(VERSION)"

