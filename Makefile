GOLANGCI_LINT_CACHE?=/tmp/golangci-lint-cache

.PHONY: fieldalignment
fieldalignment:
	-@find . -type d -not -name ".*"|while read path; do if [ `find $$path -maxdepth 1 -name *.go|wc -l` -gt 0 ]; then fieldalignment -fix `find $$path -maxdepth 1 -name *.go`; fi; done

.PHONY: formattag
formattag:
	-@find . -name "*.go" -type f -exec formattag -file {} \;

.PHONY: linter
linter: _golangci-lint-rm-unformatted-report

.PHONY: _golangci-lint-reports-mkdir
_golangci-lint-reports-mkdir:
	mkdir -p ./golangci-lint

.PHONY: _golangci-lint-run
_golangci-lint-run: _golangci-lint-reports-mkdir
	-docker run --rm \
    -v $(shell pwd):/app \
    -v $(GOLANGCI_LINT_CACHE):/root/.cache \
    -w /app \
    golangci/golangci-lint:v1.57.2 \
        golangci-lint run \
            -c .golangci.yml \
	> ./golangci-lint/report-unformatted.json

.PHONY: _golangci-lint-format-report
_golangci-lint-format-report: _golangci-lint-run
	cat ./golangci-lint/report-unformatted.json | jq > ./golangci-lint/report.json

.PHONY: _golangci-lint-rm-unformatted-report
_golangci-lint-rm-unformatted-report: _golangci-lint-format-report
	rm ./golangci-lint/report-unformatted.json

.PHONY: golangci-lint-clean
golangci-lint-clean:
	sudo rm -rf ./golangci-lint

.PHONY: tests
test:
	go test ./...