package utils

func IMax(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func IMin(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func IAbs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
