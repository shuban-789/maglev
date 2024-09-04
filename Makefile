# Build Alternative: make ARCH=$GOARCH OS=$GOOS
# (e.g.) make ARCH=arm64 OS=darwin
# (e.g.) make ARCH=arm64 OS=linux

ARCH=amd64
OS=linux

all: build

build:
  GOARCH=$(ARCH) GOOS=$(OS) go build -o maglev src/maglev.go

run:
  go run src/maglev.go

clean:
  rm -f maglev
