package baila

// Main Temperature
type Main struct {
	Temp float32
}

// Temperature type
type Temperature struct {
	Main Main
}

// TemperatureServicer temperature
type TemperatureServicer interface {
	Temperature(city string, unit string) (Temperature, error)
}
