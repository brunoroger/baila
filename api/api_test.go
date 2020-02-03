package api

import (
	"testing"
)

func TestApiCelsious(t *testing.T) {
	obj1, errTest1 := NewTemperatureServicer().Temperature("joinville", "1")

	if obj1 == nil && errTest1 != nil {
		t.Error("the temperature cannot be equal to null")
	}
}

func TestApiFahrenheit(t *testing.T) {
	obj2, errTest2 := NewTemperatureServicer().Temperature("joinville", "2")

	if obj2 == nil && errTest2 != nil {
		t.Error("the temperature cannot be equal to null")
	}
}

func TestApiUnitInvalid(t *testing.T) {
	obj3, errTest3 := NewTemperatureServicer().Temperature("joinville", "3")

	if obj3 != nil && errTest3 == nil {
		t.Error("the unit cannot be found")
	}
}
