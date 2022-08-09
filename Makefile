test:
	@./go.test.sh
.PHONY: test

coverage:
	@./go.coverage.sh
.PHONY: coverage

build:
	CGO_ENABLED=0 go build ./cmd/mvthbot

clear:
	rm mvthbot

check_generated: generate
	git diff --exit-code
.PHONY: check_generated
