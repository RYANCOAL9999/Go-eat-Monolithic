// Package calendar implements a scheduler for genertic types of things.

package calendar

// GetRota returns the calendar rota for the day filter by team
func GetRota(team string) []string {
	switch team {
	case "kitchen":
		return []string{"PersonA", "PersonB", "PersonC"}
	}

	return []string{}
}
