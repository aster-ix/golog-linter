.PHONY: all plugin exec test clean run help

all: plugin

plugin: 
	mkdir -p linter

	echo "plugin built [+]"

exec: 
	go build -o gologlinter main.go
	echo "executable built [+]"

test:
	go test ./analyzer
	echo "test finished [+]"

clean:
	rm -rf gologlinter gologlinter.so
	echo "cleaned [+]"

run: plugin
	golangci-lint run --verbose testdata/tests.go

help:
	echo "plugin	- build as plugin (mac/linux only)"
	echo "exec		- build via executable"
	echo "test 		- run tests"
	echo "clean		- remove build files"
	ecgo "run		- build plugin and run golangci-lint"