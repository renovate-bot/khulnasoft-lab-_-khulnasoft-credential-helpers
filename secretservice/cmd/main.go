//go:build linux && cgo

package main

import (
	"github.com/khulnasoft-lab/khulnasoft-credential-helpers/credentials"
	"github.com/khulnasoft-lab/khulnasoft-credential-helpers/secretservice"
)

func main() {
	credentials.Serve(secretservice.Secretservice{})
}
