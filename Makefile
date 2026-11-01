.PHONY: test lint demo logs

RUN = go run ./plugin/scripts

test:
	test -z "$$(gofmt -l plugin tests examples)"
	go vet ./...
	cd examples/programs && go vet ./...
	go test -race ./...

lint:
	claude plugin validate ./plugin --strict
	claude plugin validate . --strict

demo:
	$(RUN) review examples/logs/shop-v1.log
	$(RUN) compare examples/logs/shop-v1.log examples/logs/shop-v2.log

# Rebuilds the shop logs from the racy programs in examples/programs.
# Races are timing dependent, so counts can change between runs.
# go test exits 1 when it finds races, which is expected here.
logs:
	cd examples/programs && go test -race -count=1 ./inventory/ ./healthy/ ./cmd/... > ../logs/shop-v1.log 2>&1 || true
	cd examples/programs && go test -race -count=1 -v ./inventory/ ./healthy/ ./cmd/... > ../logs/shop-v1-verbose.log 2>&1 || true
	cd examples/programs && go test -race -count=1 -json ./inventory/ ./healthy/ ./cmd/... > ../logs/shop-v1.json 2>&1 || true
	cd examples/programs && go test -race -count=1 -tags v2 ./inventory/ ./healthy/ ./cmd/... > ../logs/shop-v2.log 2>&1 || true
	cd examples/programs && go test -race -count=1 ./healthy/ > ../logs/healthy.log 2>&1
	cd examples/programs && go run -race ./cmd/warm > ../logs/warm.log 2>&1 || true
	cd examples/programs && go test -race -count=1 ./testonly/ > ../logs/testonly.log 2>&1 || true
