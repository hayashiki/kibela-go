export GO111MODULE ?= on

VERSION  := $(shell cat VERSION)
REVISION := $(shell git rev-parse --short HEAD)

.PHONY: tag
tag:
	git tag -a $(VERSION) -m "Release $(VERSION)"
	git push --tags
