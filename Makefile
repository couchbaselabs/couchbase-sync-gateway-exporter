start:
	@echo "Ensure network exists..."
	@docker network create workshop -d bridge || true
	@echo "Running cb-server..."
	@docker run -d --name cb-server --network workshop -p 8091-8094:8091-8094 -p 11210:11210 connectsv/server:6.0.1-enterprise
	@echo "Running sync-gateway-travel_2.5..."
	@docker run -p 4984-4985:4984-4985 --network workshop --name sync-gateway-travel -d -v `pwd`/sync-gateway-config.json:/etc/sync_gateway/sync_gateway.json connectsv/sync-gateway:2.5.0-enterprise  -adminInterface :4985 /etc/sync_gateway/sync_gateway.json
.PHONY: start

stop:
	@echo "Stopping sync-gateway-travel_2.5 and cb-server  ..."
	@docker rm --force sync-gateway-travel cb-server || true

rm:
	@echo "Deleting sync-gateway-travel_2.5 and cb-server  ..."
	@docker stop sync-gateway-travel cb-server || true
	@docker network rm workshop || true
.PHONY: rm

open:
	@open http://localhost:8091
	@open http://localhost:4985/_expvar
.PHONY: open

lite:
	@./.bin/cblite cp --bidi --continuous --user demo:password ws://localhost:4984/travel-sample/ travel.cblite2
.PHONY: lite
