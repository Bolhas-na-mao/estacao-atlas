.PHONY: run build test snapshot clean release

run:
	go run .

build:
	go build ./...

test:
	go test ./...

snapshot:
	goreleaser release --snapshot --clean

clean:
	rm -rf dist/

release:
	@test -n "$(v)" || (echo "Usage: make release v=0.1.0"; exit 1)
	git tag v$(v)
	git push origin v$(v)
