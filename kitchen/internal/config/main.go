// Package config enables env vars to be used easily within code.
package config

import (
	"os"
	"strconv"
)

var defaultValues = map[string]interface{}{
	"VERSION":         "unset",
	"FSVAULT_DATADIR": "/tmp/",
}

// String rerurns a string value from the env or default or empty string
func String(key string) string {

	if defaultValue, ok := defaultValues[key].(string); ok {
		return getEnvVar(key, defaultValue).(string)
	}
	return ""
}

// Int returns an int value from the env or default or -1
func Int(key string) int {

	if defaultValue, ok := defaultValues[key].(int); ok {
		return getEnvVar(key, defaultValue).(int)
	}
	return -1
}

// Bool returns a bool value from the env or default or false
func Bool(key string) bool {

	if defaultValue, ok := defaultValues[key].(bool); ok {
		return getEnvVar(key, defaultValue).(bool)
	}
	return false
}

// getEnvVar returns the env var matching the key, or the fallback value,
func getEnvVar(key string, fallback interface{}) interface{} {

	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}

	switch fallback.(type) {
	case string:
		return value
	case bool:
		valueAsBool, err := strconv.ParseBool(value)
		if err != nil {
			return fallback
		}
		return valueAsBool
	case int:
		valueAsInt, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}
		return valueAsInt
	}
	return fallback
}
