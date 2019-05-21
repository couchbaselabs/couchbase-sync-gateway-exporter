start:
	docker-compose build
	docker-compose up
.PHONY: start

rm:
	yes | docker-compose rm
.PHONY: rm

open:
	@open http://localhost:8091
	@open http://localhost:4985/_expvar
.PHONY: open

lite:
	@./.bin/cblite cp --bidi --continuous --user demo:password ws://localhost:4984/travel-sample/ travel.cblite2
.PHONY: lite

lint:
	golangci-lint run --enable-all --deadline 5m ./...
.PHONY: lint

test:
	go test -v -failfast -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.txt ./... -timeout=2m
.PHONY: test

cover: test
	go tool cover -html=coverage.txt
.PHONY: cover
