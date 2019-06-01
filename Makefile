start:
	docker-compose build
	docker-compose up -d
.PHONY: start

stop:
	docker-compose down
.PHONY: start

lite:
	@./.bin/cblite cp --bidi --continuous --user demo:password ws://localhost:4984/travel-sample/ travel.cblite2
.PHONY: lite

lint:
	golangci-lint run --enable-all --disable dupl --deadline 5m ./...
.PHONY: lint

test:
	go test -v -failfast -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.txt ./... -timeout=2m
.PHONY: test

ci: test lint
.PHONY: ci

cover: test
	go tool cover -html=coverage.txt
.PHONY: cover

grafana:
	@echo "Generating dashboard at ./grafana/dashboard.json"
	@jsonnet -J grafana grafana/dashboard.jsonnet -o ./grafana/dashboard.json
.PHONY: grafana

grafana-dev:
	@make grafana
	@./scripts/setup-grafana
.PHONY: grafana-dev
