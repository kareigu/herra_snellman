build:
	go build -v -o bin/main src/main.go

run:
	go run src/main.go

clean:
	go clean

docker:
	docker build -t mxr/herra_snellman:latest .
