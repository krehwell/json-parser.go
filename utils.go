package main

func isValidNumber(number string) bool {
	if number[0] == '-' {
		if len(number) < 1 {
			return false
		}

		if number[1] == '0' {
			return false
		}
	}

	if number[0] == '0' {
		return false
	}

	return true
}
