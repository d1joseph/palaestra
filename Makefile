BINARY_NAME=palaestra

build:
	go build -o ${BINARY_NAME}.exe cmd/api/main.go

run: build
	./${BINARY_NAME}.exe

clean:
	go clean
	rm ${BINARY_NAME}-win