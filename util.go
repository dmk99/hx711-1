package hx711

// CalculateCalibratedReading calculate the calibrated reading based on the raw value and tare/calibrated values
func CalculateCalibratedReading(rawValue int32, attribute *HX711Attributes) (float64, error) {
	reading := float64(rawValue) - attribute.Tare
	calibrated := float64(attribute.CalibratedReading - attribute.Tare)

	ratio := float64(calibrated) / attribute.CalibratedWeight

	return ratio * reading, nil
}