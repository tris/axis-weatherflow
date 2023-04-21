package main

import (
	"math"
)

func degreesToCardinal(degrees float64) string {
	// Convert degrees to radians
	radians := math.Pi * degrees / 180.0

	// Define the cardinal directions
	directions := []string{"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE", "S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW"}

	// Calculate the index of the closest cardinal direction
	index := int(math.Round(radians / (2*math.Pi / 16.0)))
	index %= 16

	// Return the closest cardinal direction
	return directions[index]
}

// Helper functions to create pointers to basic types
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func boolPtr(b bool) *bool {
	return &b
}
