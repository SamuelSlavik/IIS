package validators

import (
	"regexp"

	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
)

type ValidatorErr struct {
	Name string
	Desc string
}

func PasswordMatch(pwd1 string, pwd2 string, validator_errs *[]ValidatorErr) {
	if pwd1 != pwd2 {
		*validator_errs = append(*validator_errs, ValidatorErr{
			Name: "PasswordMatchErr",
			Desc: "Passwords do not match!",
		})
	}
}

func EmailValidator(email string, validator_errs *[]ValidatorErr) {
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

func RoleValidator(role string, validator_errs *[]ValidatorErr) {
	switch role {
	case string(models.AdminRole), string(models.SuperuserRole), string(models.DispatcherRole), string(models.TechnicianRole), string(models.DriverRole):
		return
	}

	*validator_errs = append(*validator_errs, ValidatorErr{
		Name: "RoleErr",
		Desc: "Invalid role!",
	})
}

func HasRoleValidator(user_id uint, validator_errs *[]ValidatorErr, permitted_roles ...models.Role) {
	var user models.User

	if result := utils.DB.First(&user, user_id); result.Error != nil {
		*validator_errs = append(*validator_errs, ValidatorErr{
			Name: "InvalidUserID",
			Desc: "User not found under ID",
		})
	}

	for _, permitted_role := range permitted_roles {
		if permitted_role == user.Role {
			return
		}
	}

	*validator_errs = append(*validator_errs, ValidatorErr{
		Name: "RoleErr",
		Desc: "Not a permitted role!",
	})
}
