.PHONY: all
all:

.PHONY: gen
gen:
	go generate -v ./...

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

.PHONY: run_api
run_api:
	go run app/api/cmd/main.go --config=configs/api/config.toml
