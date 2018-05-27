package hx711

import (
	"errors"
	"github.com/montanaflynn/stats"
	"github.com/mrmorphic/hwio"
)

const (
	GAIN_A_128 = 1
	GAIN_B_32  = 2
	GAIN_A_64  = 3
)

type HX711Attributes struct {
	// Tare the reading at "zero"
	Tare float64 `json:"tare"`

	// CalibratedReading the raw reading from the device for the CalibratedWeight
	CalibratedReading float64 `json:"calibratedReading"`

	// CalibratedWeight the known weight when calibrating
	CalibratedWeight float64 `json:"calibratedWeight"`
}

type HX711 struct {
	Clock      string
	Data       string
	Gain       int
	Attributes *HX711Attributes
	clkPin     hwio.Pin
	dataPin    hwio.Pin
}

//New instantiates a new object
func New(data string, clock string) (*HX711, error) {
	var err error
	var clkPin, dataPin hwio.Pin
	if clkPin, err = hwio.GetPin(clock); err != nil {
		return &HX711{}, err
	}
	if dataPin, err = hwio.GetPin(data); err != nil {
		return &HX711{}, err
	}
	if err = hwio.PinMode(dataPin, hwio.INPUT); err != nil {
		return &HX711{}, err
	}
	if err = hwio.PinMode(clkPin, hwio.OUTPUT); err != nil {
		return &HX711{}, err
	}
	return &HX711{Clock: clock, Data: data, Gain: GAIN_A_128, clkPin: clkPin, dataPin: dataPin}, err
}

// New initializes a new object with the known calibration values
func NewWithKnownAttributes(data string, clock string, attributes *HX711Attributes) (*HX711, error) {
	hx711, err := New(data, clock)
	hx711.Attributes = attributes

	if attributes.Tare == 0 || attributes.CalibratedWeight == 0 || attributes.CalibratedReading == 0 {
		return &HX711{}, errors.New("unset attributes supplied")
	}

	return hx711, err
}

//OnReady Blocks until the chip is ready to send data
func (h *HX711) OnReady() error {
	if err := h.clockLow(); err != nil {
		return err
	}
	ready := false
	for !ready {
		r, err := h.readBit()
		if err != nil {
			return err
		}
		if r == hwio.LOW {
			ready = true
		}
	}
	return nil
}

//Sleep the chip until you need
func (h *HX711) Sleep() error {
	return h.clockHigh()
}

//Reset / Wakeup  the chip
func (h *HX711) Reset() error {
	if err := h.clockHigh(); err != nil {
		return err
	}
	hwio.DelayMicroseconds(60)
	return h.clockLow()
}

//SetGain sets the gain after data has been read
func (h *HX711) SetGain() error {
	for i := 0; i < h.Gain; i++ {
		if err := h.tick(); err != nil {
			return err
		}
	}
	return nil
}

//ReadData gets a 24bit signed int from the chip
func (h *HX711) ReadData() (int32, error) {
	c := int32(0)
	if err := h.OnReady(); err != nil {
		return 0, err
	}
	for i := 0; i < 24; i++ {
		h.tick()
		b, err := hwio.DigitalRead(h.dataPin)
		if err != nil {
			return 0, err
		}
		c = c << 1
		if b == hwio.HIGH {
			c++
		}
	}

	return twosComp(c), h.SetGain()
}

// ReadCalibratedData get the values from the device and transform them per the known attributes
func (h *HX711) ReadCalibratedData() (float64, error) {
	rawReading, err := h.ReadData()

	if err != nil {
		return 0, err
	}

	return CalculateCalibratedReading(rawReading, h.Attributes)
}

// Tare get the tare value and set it for the device for the "zero" reading
func (h *HX711) Tare(numberOfReadings int) (float64, error) {
	readings, err := h.getReadings(numberOfReadings)

	if err != nil {
		return 0, err
	}

	tare, err := stats.Percentile(readings, 50)

	if err != nil {
		return 0, err
	}

	h.Attributes.Tare = tare

	return tare, nil
}

// Calibrate calibrate the device by using a known weight
func (h *HX711) Calibrate(numberOfReadings int, knownWeight float64) error {
	readings, err := h.getReadings(numberOfReadings)

	if err != nil {
		return err
	}

	calibration, err := stats.Percentile(readings, 50)

	if err != nil {
		return err
	}

	h.Attributes.CalibratedReading = calibration
	h.Attributes.CalibratedWeight = knownWeight

	return nil
}

func (h *HX711) getReadings(numberOfReadings int) ([]float64, error) {
	if numberOfReadings <= 0 {
		return []float64{}, errors.New("must have readings >= 1")
	}

	readings := make([]float64, numberOfReadings)
	for i := 0; i < numberOfReadings; i++ {
		reading, err := h.ReadData()

		if err != nil {
			return []float64{}, err
		}

		readings[i] = float64(reading)
	}

	return readings, nil
}

func (h *HX711) tick() error {
	err := h.clockHigh()
	if err != nil {
		return err
	}
	return h.clockLow()
}

func (h *HX711) clockHigh() error {
	return hwio.DigitalWrite(h.clkPin, hwio.HIGH)
}

func (h *HX711) clockLow() error {
	return hwio.DigitalWrite(h.clkPin, hwio.LOW)
}

func (h *HX711) readBit() (int, error) {
	return hwio.DigitalRead(h.dataPin)
}

func twosComp(i int32) int32 {
	if (i & 0x800000) > 0 {
		i |= ^0xffffff
	}
	return i
}
