package hx711

import (
	"math"
	"testing"
)

func TestCalculateCalibratedReadingNegativeScale(t *testing.T) {
	// Setup
	attributes := HX711Attributes{Tare: -9574.5, CalibratedReading: -16641, CalibratedWeight: 1500}

	// Execute
	calibratedReading, _ := CalculateCalibratedReading(-16641, &attributes)

	// Verify
	if math.Round(calibratedReading) != 1500 {
		t.Errorf("Expected: %v, Actual: %v, ", 1500, math.Round(calibratedReading))
		t.Fail()
	}
}

func TestCalculateCalibratedReadingPositiveScale(t *testing.T) {
	// Setup
	attributes := HX711Attributes{Tare: -2607, CalibratedReading: -4922, CalibratedWeight: 1500}

	// Execute
	calibratedReading, _ := CalculateCalibratedReading(-4922, &attributes)

	// Verify
	if math.Round(calibratedReading) != 1500 {
		t.Errorf("Expected: %v, Actual: %v, ", 1500, math.Round(calibratedReading))
		t.Fail()
	}
}

func TestCalculateCalibratedReadingPositiveScaleBelowCalibration(t *testing.T) {
	// Setup
	attributes := HX711Attributes{Tare: -2607, CalibratedReading: -4922, CalibratedWeight: 1500}

	// Execute
	calibratedReading, _ := CalculateCalibratedReading(-4150, &attributes)

	// Verify
	if math.Round(calibratedReading) != 1000 {
		t.Errorf("Expected: %v, Actual: %v, ", 1000, math.Round(calibratedReading))
		t.Fail()
	}
}

func TestCalculateCalibratedReadingPositiveScaleAboveCalibration(t *testing.T) {
	// Setup
	attributes := HX711Attributes{Tare: -2607, CalibratedReading: -4922, CalibratedWeight: 1500}

	// Execute
	calibratedReading, _ := CalculateCalibratedReading(-7237, &attributes)

	// Verify
	if math.Round(calibratedReading) != 3000 {
		t.Errorf("Expected: %v, Actual: %v, ", 3000, math.Round(calibratedReading))
		t.Fail()
	}
}

func TestCalculateCalibratedReadingZero(t *testing.T) {
	// Setup
	attributes := HX711Attributes{Tare: -2607, CalibratedReading: -4922, CalibratedWeight: 1500}

	// Execute
	calibratedReading, _ := CalculateCalibratedReading(-2607, &attributes)

	// Verify
	if math.Round(calibratedReading) != 0 {
		t.Errorf("Expected: %v, Actual: %v, ", 0, math.Round(calibratedReading))
		t.Fail()
	}
}
