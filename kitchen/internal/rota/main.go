// Package rota implements the workday schedule for this team.

package rota

import (
	"goeat/staff"
)

func fetchRota() []string {

	rota := staff.GetKitchenRota()

	// do something with the rota

	return rota
}
