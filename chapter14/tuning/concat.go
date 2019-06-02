package tuning

func concat(vals ...string) string {
	finalVal := ""
	for i := 0; i < len(vals); i++ {
		finalVal += vals[i]
		if i != len(vals)-1 {
			finalVal += " "
		}
	}
	return finalVal
}
