
BINARY_NAME = main

build:
	go build -o ${BINARY_NAME} main.go

run:
	GIN_MODE=release ./${BINARY_NAME}

clean:
	rm ${BINARY_NAME}