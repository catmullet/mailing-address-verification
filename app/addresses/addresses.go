package addresses

import (
	"regexp"
	"strings"
)

func CheckMailingAddress(mAddr, country string) bool {

	match := true

	regex, mustValidate, _ := GetCountryRegex(country)

	if mustValidate {
		match, _ = regexp.MatchString(regex, strings.ToLower(mAddr))
	}

	return !mustValidate || match

}