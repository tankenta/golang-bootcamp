package lenconv

func MToFt(m Metre) Feet {
	return Feet(m * 3.2808)
}

func FtToM(ft Feet) Metre {
	return Metre(ft / 3.2808)
}
