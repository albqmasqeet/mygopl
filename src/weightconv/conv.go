package weightconv

func KtoL(k KiloGram) Pound {
	return Pound(k * 2.20462)
}

func LtoK(lb Pound) KiloGram {
	return KiloGram(lb * 0.453592)
}
