ROOT := $(shell pwd)
DIST_DIR := dist
OSX_APP := jhr.app
OSX_APP_FOLDER := $(OSX_APP)/Contents/MacOS

all:
	$(MAKE) build-osx-app
	$(MAKE) build-win-app

.PHONY: build-client
build-client:
	cd client && npm run build

build-osx-app: build-client
	mkdir -p $(DIST_DIR)/$(OSX_APP_FOLDER)
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin \
	cd server/cmd/jhr-server && go build -ldflags="-w -s" -o $(ROOT)/$(DIST_DIR)/$(OSX_APP_FOLDER)/jhr

build-win-app: build-client
# should we add `-H windowsgui` to -ldflags?
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows \
	cd server/cmd/jhr-server && go build -ldflags="-w -s" -o $(ROOT)/$(DIST_DIR)/jhr.exe

open-osx-app:
	open $(DIST_DIR)/$(OSX_APP)

