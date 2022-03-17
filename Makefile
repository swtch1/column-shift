export SHELL:=/bin/bash
export SHELLOPTS:=$(if $(SHELLOPTS),$(SHELLOPTS):)pipefail:errexit

CWD=$(shell pwd)

.PHONY: all build

all: build

build:
	# GOOS=darwin GOARCH=amd64 go build -o build/silver-lining-etl-macos .
	GOOS=windows GOARCH=amd64 go build -o build/silver-lining-etl-windows.exe .

package: build
	mkdir -p build/silver-lining-etl
	# cd build; cp silver-lining-etl-macos silver-lining-etl
	cd build; cp silver-lining-etl-windows.exe silver-lining-etl
	cd build; zip -r silver-lining-etl.zip silver-lining-etl

clean:
	rm -rf build

