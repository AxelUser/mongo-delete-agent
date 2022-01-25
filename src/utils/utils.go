package utils

func Contains(l []string, t string) bool {
	for _, v := range l {
		if v == t {
			return true
		}
	}

	return false
}
