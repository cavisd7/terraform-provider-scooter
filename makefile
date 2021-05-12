TAG_COMMIT := $(shell git rev-list --abbrev-commit --tags --max-count=1)
TAG := $(shell git describe --abbrev=0 --tags ${TAG_COMMIT} 2>/dev/null || true)
VERSION := $(TAG:v%=%)
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
NAME := terraform-provider-scooter_v$(VERSION)

default: build

build:
	mkdir -p ~/.terraform.d/plugins/nil.xyz/ns/scooter/$(VERSION)/$(GOOS)_$(GOARCH)
	go build -o ~/.terraform.d/plugins/nil.xyz/ns/scooter/$(VERSION)/$(GOOS)_$(GOARCH)/$(NAME)
	go run api/main.go
