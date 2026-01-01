package main

func isValidLength(password string) bool {
	n := len(password)
	if n >= 5 && n <= 12 {
		return true
	}
	return false
}
func isUppercase(char rune) bool {
	if char >= 'A' && char <= 'Z' {
		return true
	}
	return false
}
func isDigit(digit rune) bool {
	if digit >= '0' && digit <= '9' {
		return true
	}
	return false
}
func isValidPassword(password string) bool {
	// ?
	if !isValidLength(password) {
		return false
	}
	hasDigit := false
	hasUpper := false
	for _, char := range password {
		if isUppercase(char) {
			hasUpper = true
		}
		if isDigit(char) {
			hasDigit = true
		}
	}
	if !hasDigit || !hasUpper {
		return false
	}
	return true
}
