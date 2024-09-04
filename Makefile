ARCH=amd64
OS=linux

all:
  GOARCH=$(ARCH) GOOS=$(OS) go build -o maglev src/maglev.go

build:
  GOARCH=$(ARCH) GOOS=$(OS) go build -o maglev src/maglev.go

run:
  go run src/maglev.go

clean:
  rm -f maglev
