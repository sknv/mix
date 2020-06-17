.PHONY: all
all:

.PHONY: gen
gen:
	go generate -v ./... && scripts/protoc.sh

##
# Go modules section
##

.PHONY: update
update:
	go get -u all

.PHONY: prune
prune:
	go mod tidy

##
# Development
##

.PHONY: run_account
run_api:
	go run app/account/cmd/main.go --config=configs/account/config.toml

.PHONY: run_api
run_api:
	go run app/api/cmd/main.go --config=configs/api/config.toml
