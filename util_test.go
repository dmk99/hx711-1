package hx711

import "testing"

func TestCalculateCalibratedReading(t *testing.T) {
	// Setup
	attributes := HX711Attributes{Tare:500, CalibratedReading:1000, CalibratedWeight:3000}

	// Execute
	calibratedReading, _ := CalculateCalibratedReading(9000, &attributes)

	// Verify
	if calibratedReading != 1500 {
		t.Fail()
	}
}