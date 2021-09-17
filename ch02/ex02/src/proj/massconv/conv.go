package massconv

func KgToLb(kg KiloGram) Pound {
	return Pound(kg * 2.2046)
}

func LbToKg(lb Pound) KiloGram {
	return KiloGram(lb / 2.2046)
}
