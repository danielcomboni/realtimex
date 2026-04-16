run-example:
	go run examples/gin/main.go

test:
	go test ./...

fmt:
	go fmt ./...

tidy:
	go mod tidy