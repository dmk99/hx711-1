package hx711

// CalculateCalibratedReading calculate the calibrated reading based on the raw value and tare/calibrated values
func CalculateCalibratedReading(rawValue int32, attribute *HX711Attributes) (float64, error) {
	weight := ((float64(rawValue) - attribute.Tare) / (attribute.CalibratedReading - attribute.Tare)) * attribute.CalibratedWeight
	return weight, nil
}
