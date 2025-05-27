//go:build dev

package datastore

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {

	defer SetupTestVaultRoot()()

	orderId := time.Now().Unix()
	order := "Spaghetti Putanesca"

	// add an order
	err := StoreOrder(orderId, []byte(order))
	assert.Equal(t, nil, err, "store order err")

	// get the order
	data, err := GetOrder(orderId)
	assert.Equal(t, nil, err, "get order err")

	assert.Equal(t, order, string(data), "stored order matches get order")
}

func TestListOrders(t *testing.T) {

	defer SetupTestVaultRoot()()

	var testOrders = []struct {
		orderId int64
		order   string
	}{
		{orderId: int64(1), order: "Haggis, neeps, and tatties"},
		{orderId: int64(2), order: "Cranachan"},
		{orderId: int64(3), order: "Cullen Skink"},
	}

	// test no orders
	orders := ListOrders()
	assert.Equal(t, 0, len(orders), "no orders stored")

	// add test orders
	for _, testOrder := range testOrders {
		StoreOrder(testOrder.orderId, []byte(testOrder.order))
	}

	// test orders count
	orders = ListOrders()
	assert.Equal(t, len(testOrders), len(orders), "orders stored")

	// get the orders
	for _, testOrder := range testOrders {
		data, err := GetOrder(testOrder.orderId)
		assert.Equal(t, nil, err, "get order error")
		assert.Equal(t, testOrder.order, string(data), "get order")
	}
}
