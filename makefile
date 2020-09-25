# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=fops
BINARY_DIR=bin
BINARY_PATH=$(BINARY_DIR)/$(BINARY_NAME)
PROJ_PATH=kiedveian/JDExam/
REPO_PATH=$(PROJ_PATH)/fops
FULL_GITHUB_PROJ_PATH=github.com/$(PROJ_PATH)
FULL_GITHUB_REPO_PATH=github.com/$(REPO_PATH)

VERSION=v0.0.2
LDFLAGS=-ldflags "-X github.com/kiedveian/JDExam/fops.Version=$(VERSION)"

all: deps test build
build:
		$(GOBUILD) $(LDFLAGS) -o $(BINARY_PATH) -v $(FULL_GITHUB_PROJ_PATH)
test:
		$(GOTEST) -v $(FULL_GITHUB_REPO_PATH)
clean:
		$(GOCLEAN)
		rm -f $(BINARY_PATH)
run:
		./$(BINARY_PATH)
deps:
		$(GOGET) $(FULL_GITHUB_PROJ_PATH)
