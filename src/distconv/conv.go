package distconv

// convert feet to meter
func FtoM(f Feet) Meter {
	return Meter(f * 0.3048)
}

// convert meter to feet
func MtoF(m Meter) Feet {
	return Feet(m * 3.28084)
}
