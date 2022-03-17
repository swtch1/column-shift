export SHELL:=/bin/bash
export SHELLOPTS:=$(if $(SHELLOPTS),$(SHELLOPTS):)pipefail:errexit

CWD=$(shell pwd)

.PHONY: all build

all: build

build.%:
	GOOS=darwin GOARCH=amd64 go build -o build/silver-lining-etf-macos .
	# GOOS=windows GOARCH=amd64 go build -o build/silver-lining-etf-windows .

package: build
	mkdir -p build/silver-lining-etf
	cd build; cp silver-lining-etf-macos silver-lining-etf
	cd build; cp silver-lining-etf-windows silver-lining-etf.exe
	cd build; zip -r silver-lining-etf.zip silver-lining-etf

clean:
	rm -rf build

