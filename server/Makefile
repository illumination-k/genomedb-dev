.PHONY: fmt

fmt:
	go fmt ./...

test:
	go test `go list ./... | grep -v -e "genomedb/ent"`
