package datastore

import (
	"goeat/kitchen/internal/config"
	"log"
	"path"
)

var (

	// using a var allows simply mocking of the datastore
	getVaultRoot = func() string {
		return path.Join(config.String("FSVAULT_DATADIR"), "/kitchen/")
	}
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("kitchen.datastore.init(): root vault dir", getVaultRoot())
}
