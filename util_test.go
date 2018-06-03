package hx711

import (
	"testing"
	"math"
)

func TestCalculateCalibratedReading(t *testing.T) {
	// Setup
	attributes := HX711Attributes{Tare:5193, CalibratedReading:-2632, CalibratedWeight:1500}

	// Execute
	calibratedReading, _ := CalculateCalibratedReading(-2632, &attributes)

	// Verify
	if math.Round(calibratedReading) != 1500 {
		t.Fail()
	}
}