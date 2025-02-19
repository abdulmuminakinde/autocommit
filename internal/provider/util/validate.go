package util

import "fmt"

// ValidateTemperature validates the temperature for the AI provider APi
func ValidateTemperature(temperature float32) error {
	if temperature <= 0 || temperature > 1 {
		return fmt.Errorf("temperature must be greater than 0 and less than or equal to 1")
	}
	return nil
}
