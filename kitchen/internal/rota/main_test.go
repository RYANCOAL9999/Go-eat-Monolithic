//go:build test

package rota

import (
	"goeat/staff"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchRota(t *testing.T) {

	mockRota := []string{"A", "B"}
	mockStaffGetKitchenRota := func() []string {
		return mockRota
	}

	// override the 'api' call with our mock function
	staff.GetKitchenRota = mockStaffGetKitchenRota

	assert.Equal(t, fetchRota(), mockRota, "expect A, B")
}
