SHELL := /bin/bash

# =====================================================================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor

# =====================================================================================================================
# Running tests locally

test:
	go test -v ./... -count=1

test-heapsort:
	go test -v ./heapsort/heapsort_test.go -count=1

test-quicksort:
	go test -v ./quicksort/quicksort_test.go -count=1

test-deque:
	go test -v ./deque/deque_test.go -count=1

