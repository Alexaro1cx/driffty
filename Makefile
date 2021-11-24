TARGET_BRANCH ?= master

lint:
	golangci-lint run -c .golangci.yml
lint-ci:
	golangci-lint run -c .golangci.yml --new-from-rev=${TARGET_BRANCH} --out-format github-actions
lint-fix:
	golangci-lint run -c .golangci.yml --fix
