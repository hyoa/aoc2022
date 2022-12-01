dev:
	go run cmd/day/day.go $(ARGS)

test-one:
	go test internal/day/day_test.go -run /$(ARGS) -v

test-all:
	go test internal/day/day_test.go