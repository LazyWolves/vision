GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
BINARY_NAME=vision
DIST_DIR=dist
SRC=main/vision.go
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
		@echo "installed vision"

uninstall:
		@rm -f $(DESTDIR)$(prefix)/bin/$(BINARY_NAME)
		@echo "uinsalled vision"
