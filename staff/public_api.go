package staff

import "goeat/staff/internal/calendar"

// The public api implemented as vars to make mocking
// easier.
var (
	// GetKitchenRota returns a string slice of rota entries
	GetKitchenRota = func() []string {
		return calendar.GetRota("kitchen")
	}
)
