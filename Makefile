# https://www.gnu.org/software/make/

# Variables
##VAR    ARCH      : architecture to be used for cross compiling the azdevman binary
ARCH ?= amd64
##VAR    OS        : the operating system type for cross compiling the azdevman binary
OS   ?= linux

BIN_DIR   ?= $(dir bin/)
BIN_NAME  ?= "azdevman"

# ~~~~~
# Usage
# ~~~~~
.DEFAULT: help
.PHONY: help
help: Makefile
	@echo "Usage: make [-e VARIABLES...] [TARGETS...]"
	@echo "Variables:"
	@sed -n 's/^##VAR//p' $< | sort
	@echo "Targets:"
	@sed -n 's/^##TAR//p' $< | sort

##TAR    binary    : builds the azdevman binary locally
.PHONY: binary
binary:
	GOOS=$(OS) GOARCH=$(ARCH) go build -ldflags "$()" -o $(BIN_DIR)$(BIN_NAME)
