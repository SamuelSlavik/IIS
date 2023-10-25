package validators

import (
	"regexp"
)

type ValidatorErr struct {
	Name string
	Desc string
}

func Password_match(pwd1 string, pwd2 string, validator_errs *[]ValidatorErr) {
	if pwd1 != pwd2 {
		*validator_errs = append(*validator_errs, ValidatorErr{
			Name: "PasswordMatchErr",
			Desc: "Passwords do not match!",
		})
	}
}

func Email_validator(email string, validator_errs *[]ValidatorErr) {
	pattern := `^(?P<name>[a-zA-Z0-9.!#$%&'*+/=?^_ \x60{|}~-]+)@(?P<domain>[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*)$`

	re, err := regexp.Compile(pattern)
	if err != nil {
		*validator_errs = append(*validator_errs, ValidatorErr{
			Name: "RegexpCompErr",
			Desc: "Regex compilation error!",
		})
	}

	if !re.MatchString(email) {
		*validator_errs = append(*validator_errs, ValidatorErr{
			Name: "EmailFmtErr",
			Desc: "Wrong email format!",
		})
	}
}
