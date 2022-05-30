package meteorology

import (
	"fmt"
)

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// String method to the TemperatureUnit type returns a string representation of the temperature unit
func (t TemperatureUnit) String() string {
	tempUnits := [...]string{"°C", "°F"}
	return tempUnits[t]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// String method to the Temperature type returns a string representation of the temperature
func (t Temperature) String() string {
	return fmt.Sprintf("%d %s", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// String method to SpeedUnit returns a string representation of the speed unit
func (s SpeedUnit) String() string {
	speedUnits := [...]string{"km/h", "mph"}
	return speedUnits[s]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// String method for Speed struct returns a string representation of the speed
func (s Speed) String() string {
	return fmt.Sprintf("%d %s", s.magnitude, s.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// String method for MeteorologyData returns a string representation of the meteorology data
// For Example "Athens: 21 °C, Wind N at 16 km/h, 63% Humidity"
func (m *MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity", m.location, m.temperature, m.windDirection, m.windSpeed, m.humidity)
}
