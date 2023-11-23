package validators

import "github.com/AdamPekny/IIS/backend/models"

func StatusValidator(status string, validator_errs *[]ValidatorErr) {
	switch status {
	case string(models.PendingStatus), string(models.InProgressStatus), string(models.DoneStatus), "":
		return
	}

	*validator_errs = append(*validator_errs, ValidatorErr{
		Name: "StatusErr",
		Desc: "Invalid status!",
	})
}
