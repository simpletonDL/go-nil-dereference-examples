package main

import "regexp"

func isValidName(name string) bool {
	if len(name) <= 1 || len(name) >= 50 {
		return false
	}
	matched, _ := regexp.MatchString(`^[A-Za-z0-9]+$`, name)
	return matched
}

func isValidAge(age int) bool {
	return age > 0 && age <= 150
}
