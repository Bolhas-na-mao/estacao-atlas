.PHONY: run build test vet ci snapshot clean release

run:
	go run .

build:
	go build ./...

test:
	go test -v ./...

vet:
	go vet ./...

ci: vet test

snapshot:
	goreleaser release --snapshot --clean

clean:
	rm -rf dist/

release:
	@test -n "$(v)" || (echo "Usage: make release v=0.1.0 [m=\"release message\"]"; exit 1)
	@if [ -n "$(m)" ]; then \
		git tag -a v$(v) -m "$(m)"; \
	else \
		git tag v$(v); \
	fi
	git push origin v$(v)
