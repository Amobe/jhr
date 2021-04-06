
DIST_DIR := dist
OSX_APP := jhr.app
OSX_APP_FOLDER := $(OSX_APP)/Contents/MacOS

all:

build-app:
	mkdir -p $(DIST_DIR)/$(OSX_APP_FOLDER)
	cd server && go build -o ../$(DIST_DIR)/$(OSX_APP_FOLDER)/jhr

open:
	open $(DIST_DIR)/$(OSX_APP)