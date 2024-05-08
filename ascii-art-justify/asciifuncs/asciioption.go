package ascii

func Center(s string) string {
	res := ""
	spc := (GetSize() / 2) - GetLen(s)/2
	for spc != 0 {
		res += " "
		spc--
	}
	return res
}

func Left(s string) string {
	res := ""
	return res
}

func Right(s string) string {
	res := ""

	spc := GetSize() - GetLen(s)
	for spc != 0 {
		res += " "
		spc--
	}
	return res
}

func Justify(s string) string {
	res := ""
	spc := (GetSize() / len(s)) - GetLen(s)
	for spc != 0 {
		res += " "
		spc--
	}
	return res
}
