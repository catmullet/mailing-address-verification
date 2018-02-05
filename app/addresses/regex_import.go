package addresses

import (
	"encoding/xml"
	"io/ioutil"
	"strings"
	"os"
)

type Countries struct {
	CountriesList []Country `xml:"Country"`
}

type Country struct {
	Code string
	Name string
	Regex string
	MustValidate bool
}

func GetCountryRegex(codeOrName string) (string, bool, error) {
	xmlFile, err := ioutil.ReadFile(os.Getenv("RFILE"))
	if err != nil {
		return "", false, err
	}

	var q Countries
	err = xml.Unmarshal(xmlFile, &q)

	if err != nil {

	}

	for _, country := range q.CountriesList {
		if strings.ToLower(country.Name) == strings.ToLower(codeOrName) || strings.ToLower(country.
			Code) == strings.ToLower(codeOrName) {
			return country.Regex, country.MustValidate, nil
		}
	}

	return "", false, err
}
