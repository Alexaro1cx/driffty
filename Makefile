TARGET_BRANCH ?= master

lint:
	golangci-lint run -c .golangci.yml
lint-ci:
	golangci-lint run -c .golangci.yml --new-from-rev=${TARGET_BRANCH} --out-format github-actions
lint-fix:
	golangci-lint run -c .golangci.yml --fix

unit-test:
	$(call print-target)
	go test  ./... -v -coverprofile cover.out
update-snapshot:
	UPDATE_SNAPSHOTS=true go test  ./...
coverage: unit-test
	$(call print-target)
	@python3 ../.github/scripts/coverage/get-coverage.py cover.out
	@go tool cover -html=./cover.out -o coverage.html

gosec:
	gosec -exclude=G104 ./...

run:
	go run . scan -q ../queries -i ./mock/unified.json -o ./output -p ./mock/projects.json
e2e: run

define print-target
	@printf "Executing target: \033[36m$@\033[0m\n"
endef
