.PHONY: all test clean


test:
	go test ./test -v -race
run:
	go run main.go