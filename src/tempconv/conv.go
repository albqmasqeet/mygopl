package tempconv

// converts a Celsius temperature to Fahrenheit.
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// converts a Celsius temperature to Kelvin.
func CtoK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)

}

// converts a Fahrenheit temperature to Celsius.
func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// converts a Fahrenheit temperature to Kelvin.
func FtoK(f Fahrenheit) Kelvin {
	return Kelvin(f*5/9 + 459.67)
}

// converts a Kelvin temperature to Celsius.
func KtoC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// converts a Kelvin temperature to Fahrenheit
func KtoF(k Kelvin) Fahrenheit {
	return Fahrenheit((k-273.15)*9/5 + 32)
}
