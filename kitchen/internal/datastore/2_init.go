//go:build dev

package datastore

import (
	"log"
	"os"
)

// SetupTestVaultRoot enables simple mocking of the fsvault datastore,
// just defer a call to the return function.
func SetupTestVaultRoot() func() {

	testVaultRoot, err := os.MkdirTemp("", "test-fsvault")
	if err != nil {
		panic(err)
	}

	// override
	getVaultRoot = func() string { return testVaultRoot }
	log.Println("stravaint.datastore.init().SetupTestVaultRoot():", testVaultRoot)

	// caller defers this returned func
	return func() { os.RemoveAll(testVaultRoot) }
}
