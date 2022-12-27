.PHONY: fmt test

fmt:
	go fmt ./...

test:
	go test `go list ./... | grep -v -e "genomedb/ent"`

prepare-files:
	echo 'todo'

test-full: prepare-files
	echo 'todo'