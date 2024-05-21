BINARY=nghtbrd
.DEFAULT_GOAL := run

run:
	/usr/lib/go/bin/go run .

install:
	go build -o nghtbrd
	mv $(BINARY) /usr/local/bin
	echo "Installation complete"

build:
	go build
