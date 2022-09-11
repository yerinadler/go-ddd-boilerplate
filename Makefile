cleanup:
	rm -rf bin
build:
	GOARCH=amd64 GOOS=linux go build -o bin/server cmd/server/main.go
run:
	./bin/server