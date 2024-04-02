//go:build darwin && cgo

package main

import (
	"github.com/khulnasoft-lab/khulnasoft-credential-helpers/credentials"
	"github.com/khulnasoft-lab/khulnasoft-credential-helpers/osxkeychain"
)

func main() {
	credentials.Serve(osxkeychain.Osxkeychain{})
}
