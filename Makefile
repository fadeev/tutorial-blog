PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=NewApp \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=blogd \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=blogcli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

all: install

install: go.sum
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/blogd
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/blogcli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

init-pre:
	rm -rf ~/.blogcli
	rm -rf ~/.blogd
	blogd init mynode --chain-id blog
	blogcli config keyring-backend test

init-user1:
	blogcli keys add user1 --output json 2>&1

init-user2:
	blogcli keys add user2 --output json 2>&1

init-post:
	blogd add-genesis-account $$(blogcli keys show user1 -a) 1000token,100000000stake
	blogd add-genesis-account $$(blogcli keys show user2 -a) 500token
	blogcli config chain-id blog
	blogcli config output json
	blogcli config indent true
	blogcli config trust-node true
	blogd gentx --name user1 --keyring-backend test
	blogd collect-gentxs

init: init-pre init-user1 init-user2 init-post