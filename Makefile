
DIST_DIR := dist
OSX_APP := jhr.app
OSX_APP_FOLDER := $(OSX_APP)/Contents/MacOS

all:

build-server:
	cd server && go build -o ../$(DIST_DIR)/$(OSX_APP_FOLDER)/jhr

build-client:
	cd client && npm run build

build-app:
	mkdir -p $(DIST_DIR)/$(OSX_APP_FOLDER)
	$(MAKE) build-client
	$(MAKE) build-server

open:
	open $(DIST_DIR)/$(OSX_APP)

