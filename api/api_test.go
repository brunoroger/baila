package api

import (
	"testing"
)

func TestApi(t *testing.T) {
	// test in celsious
	obj1, errTest1 := TemperatureServicer.Temperature(TemperatureServicer{}, "Joinville", "1")

	if obj1 == nil && errTest1 != nil {
		t.Error("the temperature cannot be equal to null")
	}

	// test in fahrenheit
	obj2, errTest2 := TemperatureServicer.Temperature(TemperatureServicer{}, "Joinville", "2")

	if obj2 == nil && errTest2 != nil {
		t.Error("the temperature cannot be equal to null")
	}

	// test unit not found
	obj3, errTest3 := TemperatureServicer.Temperature(TemperatureServicer{}, "Joinville", "3")

	if obj3 != nil && errTest3 == nil {
		t.Error("the unit cannot be found")
	}
}
