.PHONY: all plugin exec test clean run help

all: plugin

plugin: 
	golangci-lint custom -v
	echo "plugin built [+]"

exec: 
	go build -o gologlinter main.go
	echo "executable built [+]"


test:
	go test ./analyzer -verbose

clean:
	rm -rf gologlinter gologlinter.exe
	echo "cleaned [+]"

run: plugin
	cd testdata
	../gologlinter run tests.go
	cd ..
	echo "test finished [+]"

help:
	@echo "Makefile commands:"
	@echo "  plugin  - build custom golangci-lint binary"
	@echo "  exec    - build executable"
	@echo "  test    - run tests in analyzer package"
	@echo "  clean   - remove build files"
	@echo "  run     - build plugin and run golangci-lint on testdata"