# Build Alternatives: 
# make ARCH=$GOARCH OS=$GOOS
# make ARCH=arm64 OS=darwin
# make ARCH=amd64 OS=darwin
# make ARCH=arm64 OS=linux
# make ARCH=amd64 OS=linux
# make ARCH=i386 OS=windows
# make ARCH=amd64 OS=windows

ARCH=amd64
OS=linux

all: build

build:
  GOARCH=$(ARCH) GOOS=$(OS) go build -o maglev src/maglev.go

run:
  go run src/maglev.go

clean:
  rm -f maglev
