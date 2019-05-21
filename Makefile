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
