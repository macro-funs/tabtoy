VERSION := 3.1.4
BUILD_SOURCE_PACKAGE := github.com/macro-funs/tabkit/build
BINARY_PACKAGE := github.com/macro-funs/tabkit
BINARY_NAME := tabtoy
BUILD_TIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

GIT_COMMIT := $(shell git rev-parse HEAD)
VERSION_STRING := -X "$(BUILD_SOURCE_PACKAGE).BuildTime=$(BUILD_TIME)" \
                 -X "$(BUILD_SOURCE_PACKAGE).Version=$(VERSION)" \
                 -X "$(BUILD_SOURCE_PACKAGE).GitCommit=$(GIT_COMMIT)"

# 定义支持的操作系统和架构
PLATFORMS := darwin-arm64 windows-amd64 windows-386 linux-amd64 linux-386
ifeq ($(OS),Windows_NT)
	PLATFORMS := $(PLATFORMS) windows-amd64
	OS := windows
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Darwin)
		PLATFORMS := $(PLATFORMS) darwin-amd64
		OS := darwin
	endif
	UNAME_M := $(shell uname -m)
	ifeq ($(UNAME_M),x86_64)
		ARCH := amd64
		ifeq ($(UNAME_M),amd64)
			ARCH := amd64
		endif
	endif
endif

.PHONY: all $(PLATFORMS) clean test

all: $(PLATFORMS)

$(PLATFORMS):
	@mkdir -p bin/$@
	GOOS=$(OS) GOARCH=$(ARCH) go build -v -p 4 -o bin/$@/$(BINARY_NAME)$(if $(findstring windows,$(OS)),.exe) \
		-ldflags "$(VERSION_STRING)" $(BINARY_PACKAGE)
	cd bin/$@ && tar zcvf ../../$(BINARY_NAME)-$(VERSION)-$@.tar.gz $(BINARY_NAME)$(if $(findstring windows,$(OS)),.exe)

test:
	go test -v $(BINARY_PACKAGE)

clean:
	rm -rf bin
	rm -f $(BINARY_NAME)-$(VERSION)-*.tar.gz

show:
	echo $(PLATFORMS)-$(VERSION)-$(GIT_COMMIT)-$(BUILD_TIME)-$(OS)-$(ARCH)