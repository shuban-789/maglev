ARCH=amd64
OS=linux

all: build

build:
  @echo "Converting src/maglev.go --> maglev (ELF)"
  GOARCH=$(ARCH) GOOS=$(OS) go build -o maglev src/maglev.go

run:
  @echo "Running src/maglev.go..."
  go run src/maglev.go

clean:
  @echo "Removing maglev (ELF)..."
  rm -f maglev
