# Reading from an hx711 24-Bit ADC on a Raspberry Pi in Golang

This is based off [https://github.com/rajmaniar/hx711](https://github.com/rajmaniar/hx711) it extends the implementation
by adding the ability to tare and calibrate the device. It's non-opinionated about the units that you are converting
e.g. grams, pounds etc.

The `HX711Attributes` struct is introduced to handle keeping track of known device attribute data.

Both the tare and calibration use the median of the `numberOfReadings` retrieved because there are situations where a reading
from the load cell can be a large spike which could throw off an average.

To tare the device you can use `Tare(numberOfReadings int)`.
To calibrate the device you can use `Calibrate(numberOfReadings int, knownWeight float64)`.

### Example:

```
import "github.com/rajmaniar/hx711"

func main() {
    
    clock := "gpio23"
    data := "gpio24"
    
    // existing attributes are known
    attributes := &HX711.Attributes{Tare: 1000, CalibratedReading: 2000, CalibratedWeight: 1500}
    
    h,err := hx711.NewWithKnownAttributes(data, clock)
    
    if err != nil {
        fmt.Printf("Error: %v",err)
    }
    
    for err == nil {
        var data int32
    
        // returns the reading based on the calibration attributes
        data, err = h.ReadCalibratedData()
        fmt.Printf("Read from HX711: %v\n",data)
        time.Sleep(250 * time.Millisecond)
    }
    fmt.Printf("Stopped reading because of: %v\n",err)
}

```

### NB
* `h.Reset()` will reset the chip
* `h.Gain` is set to `hx711.GAIN_A_128` by default