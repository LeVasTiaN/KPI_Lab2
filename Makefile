default: out/example

clean:
	rm -rf out

test:
	go test ./...

out/example: *.go cmd/example/*.go
	mkdir -p out
	go build -o out/example ./cmd/example
