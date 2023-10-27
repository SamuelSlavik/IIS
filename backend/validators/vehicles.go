package validators

import (
	"regexp"
)

func Registration_validator(registration string, validator_errs *[]ValidatorErr) {
	pattern := "^[A-Z0-9]{3}[0-9]{4}$"

	re, err := regexp.Compile(pattern)
	if err != nil {
		*validator_errs = append(*validator_errs, ValidatorErr{
			Name: "RegexpCompErr",
			Desc: "Regex compilation error!",
		})
	}
	if !re.MatchString(registration) {
		*validator_errs = append(*validator_errs, ValidatorErr{
			Name: "RegistrationFmtErr",
			Desc: "Wrong registration number format!",
		})
	}
}
