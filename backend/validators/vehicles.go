package validators

import (
	"regexp"

	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
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

func Vehicle_type_validator(vehicle_type string, validator_errs *[]ValidatorErr) {
	res := utils.DB.Where("type = ?", vehicle_type).Find(&models.VehicleType{})
	if res.Error != nil {
		*validator_errs = append(*validator_errs, ValidatorErr{
			Name: "DatabaseErr",
			Desc: res.Error.Error(),
		})
		return
	}
	if res.RowsAffected == 0 {
		*validator_errs = append(*validator_errs, ValidatorErr{
			Name: "VehicleTypeErr",
			Desc: "Vehicle type does not exist!",
		})
	}
}
