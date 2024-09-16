build:
	go build -o bin/tasker
run: build
	./bin/tasker
test:
	go test -v ./...
