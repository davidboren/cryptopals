.PHONY: test test-coverage dep

SRC_DIR := pkg

dep:
	CC="clang" && CXX="clang++" dep ensure

test:
	go test -v ./pkg/...

test-coverage:
	go test -v -coverprofile=coverage.out ./pkg/...
