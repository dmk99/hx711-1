package hx711

// calculateScale calculate the scale. If tare or calibrated reading are zero return 1.
func calculateScale(attributes *HX711Attributes) float64 {
	if attributes.Tare != 0 && attributes.CalibratedReading != 0 {
		return (attributes.CalibratedWeight - attributes.Tare) / attributes.CalibratedReading
	}

	return 1
}

// CalculateCalibratedReading calculate the calibrated reading based on the raw value and tare/calibrated values
func CalculateCalibratedReading(rawValue int32, attribute *HX711Attributes) (float64, error) {
	weight := attribute.scale*float64(rawValue) + attribute.Tare
	return weight, nil
}
