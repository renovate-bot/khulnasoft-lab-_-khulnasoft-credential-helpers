package main

import (
	"github.com/khulnasoft-lab/khulnasoft-credential-helpers/credentials"
	"github.com/khulnasoft-lab/khulnasoft-credential-helpers/pass"
)

func main() {
	credentials.Serve(pass.Pass{})
}
