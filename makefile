# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=fops
BINARY_DIR=bin
BINARY_PATH=$(BINARY_DIR)/$(BINARY_NAME)
REPO_PATH=kiedveian/JDExam/fops
FULL_GITHUB_REPO_PATH=github.com/$(REPO_PATH)

VERSION=v0.0.1

all: test build
build:
		$(GOBUILD) -ldflags "-X main.Version=$(VERSION)" -o $(BINARY_PATH) -v $(FULL_GITHUB_REPO_PATH)
test:
		$(GOTEST) -v $(FULL_GITHUB_REPO_PATH)
clean:
		$(GOCLEAN)
		rm -f $(BINARY_PATH)
run:
		./$(BINARY_NAME)
deps:
		$(GOGET) $(FULL_GITHUB_REPO_PATH)
