all: test

deps:
	go get -d -v ./...

updatedeps:
	go get -d -v -u -f ./...

testdeps:
	go get -d -v -t ./...

updatetestdeps:
	go get -d -v -t -u -f ./...

build: deps
	go build ./...

install:
	go install ./...

lint: testdeps
	go get -v github.com/golang/lint/golint
	for file in $$(find . -name '*.go' | grep -v '\.pb.go$$' | grep -v '\.pb.log.go$$'); do \
		golint $$file; \
	done

vet: testdeps
	go vet ./...

errcheck: testdeps
	go get -v github.com/kisielk/errcheck
	errcheck ./...

pretest: lint vet errcheck

test: testdeps pretest
	go test -test.v .

clean:
	go clean -i ./...

proto: install
	go get -v go.pedge.io/protoeasy/cmd/protoeasy
	go get -v go.pedge.io/pkg/cmd/strip-package-comments
	go get -v github.com/golang/protobuf/protoc-gen-go/...
	protoeasy --go --go-import-path github.com/sr/grpcinstrument .
	find . -name *\.pb\*\.go | xargs strip-package-comments
	protoc --grpcinstrument_out=. example/example.proto

.PHONY: \
	all \
	deps \
	updatedeps \
	testdeps \
	updatetestdeps \
	build \
	lint \
	vet \
	errcheck \
	pretest \
	test \
	clean \
	proto
