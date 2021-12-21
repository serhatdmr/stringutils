package stringutils

func Join(lst []string, char string) string {

	if lst == nil {
		return ""
	}

	var res string

	for i, v := range lst {
		if i == 0 {
			res += v
		} else {
			res += char + " " + v
		}
	}
	return res
}
