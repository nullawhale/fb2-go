OUT := build/fb2-reader
VERSION := 0.0.1

all: run

server:
	go build -i -v -o ${OUT}-${VERSION} -ldflags="-X main.version=${VERSION}"

run: server
	./${OUT}-${VERSION} grav.fb2

clean:
	-@rm ${OUT}-${VERSION} ${OUT}-v*