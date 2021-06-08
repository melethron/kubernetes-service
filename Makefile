.PHONY: .build

build:
	go build -v ./cmd/kubernetes/

.DEFAULT-GOAL:= build