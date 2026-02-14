.PHONY: all plugin exec test clean run help

all: plugin

plugin: 
	go mod tidy
	go mod tidy -C testdata
	
	golangci-lint custom -v
	echo "plugin built [+]"

exec: 
	go build -o gologlinter main.go
	echo "executable built [+]"


test:
	go test ./analyzer 

test_w_linter:
	./gologlinter run testdata/tests.go
	echo "test finished [+]"


clean:
	rm -rf gologlinter gologlinter.exe
	echo "cleaned [+]"

run: plugin test_w_linter

help:
	@echo "Make commands:"
	@echo "  all            Build plugin (default target)"
	@echo "  plugin         Run go mod tidy (root + testdata) and build custom golangci-lint binary"
	@echo "  exec           Build standalone gologlinter executable"
	@echo "  test           Run Go tests in ./analyzer package"
	@echo "  test_w_linter  Run built gologlinter against testdata/tests.go"
	@echo "  run            Build plugin and run linter on testdata"
	@echo "  clean          Remove built binaries"

