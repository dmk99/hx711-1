package hx711

// CalculateCalibratedReading calculate the calibrated reading based on the raw value and tare/calibrated values
func CalculateCalibratedReading(rawValue int32, attribute *HX711Attributes) (float64, error) {
	y := attribute.CalibratedWeight
	b := attribute.Tare
	x := attribute.CalibratedReading
	m := (y - b) / x

	weight := m * float64(rawValue) + b

	return weight, nil
}