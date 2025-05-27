package datastore

import (
	"fmt"

	"github.com/thisdougb/go-fsvault/fsvault"
)

var (
	orderKeyFmt = "/order/%d"
)

// ListOrders returns the orderIds as a list of strings.
func ListOrders() []string {
	key := "/order/"
	return fsvault.List(getVaultRoot(), key)
}

// GetOrder returns a specific order.
func GetOrder(id int64) ([]byte, error) {
	key := fmt.Sprintf(orderKeyFmt, id)
	return fsvault.Get(getVaultRoot(), key)
}

// StoreOrder puts a specific order into the datastore.
func StoreOrder(id int64, data []byte) error {
	key := fmt.Sprintf(orderKeyFmt, id)
	return fsvault.Put(getVaultRoot(), key, data)
}
