EXECUTABLE=crud-go

all: run

build:
	go build .

run:
	go run .

both: ${EXECUTABLE}
	go build . && ./${EXECUTABLE}
