build:
	go mod tidy
run:
	go run functions.go main.go
test:
	go test -bench=.
