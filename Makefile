GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
BINARY_NAME=vision
DIST_DIR=dist
SRC=main/vision.go
CONF_DIR=conf
prefix=/usr/local

all: build

build:	clean
		@echo "building vision..."
		@mkdir $(DIST_DIR)
		$(GOBUILD) -o $(DIST_DIR)/$(BINARY_NAME) $(SRC)
		@echo "build successful"

clean:
		$(GOCLEAN)
		@rm -rf $(DIST_DIR)
		@echo "build cleaned"

run:
		$(GORUN) $(SRC)

install:
		@install -D $(DIST_DIR)/$(BINARY_NAME) $(DESTDIR)$(prefix)/bin/$(BINARY_NAME)
		@mkdir -p $(DESTDIR)/etc/$(BINARY_NAME)
		@cp $(CONF_DIR)/config.json $(DESTDIR)/etc/vision/config.json
		@cp $(CONF_DIR)/vision.service $(DESTDIR)/etc/systemd/system/vision.service
		@echo "installed vision"

uninstall:
		@rm -f $(DESTDIR)$(prefix)/bin/$(BINARY_NAME)
		@rm -rf $(DESTDIR)/etc/vision
		@rm -rf $(DESTDIR)/etc/systemd/system/vision.service
		@echo "uinsalled vision"

.PHONY: all build clean run install uninstall
