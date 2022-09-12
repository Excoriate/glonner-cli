#!/bin/bash
#
# This scripts is intended to make goreleaser beautiful, as it's, without mixing 'sh -c' nightmare scripts

function main() {
  cd cli || exit

  go mod tidy
  go fmt ./...
  go vet ./...
  go test -v -cover -race -timeout 120s ./...
}

main "$@"
