//go:build windows

package main

import (
	"github.com/khulnasoft-lab/khulnasoft-credential-helpers/credentials"
	"github.com/khulnasoft-lab/khulnasoft-credential-helpers/wincred"
)

func main() {
	credentials.Serve(wincred.Wincred{})
}
