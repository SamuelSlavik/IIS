package validators

import (
	"time"

	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
)

func DeadlineValidator(deadline time.Time, validator_errs *[]ValidatorErr) {
	if deadline.Truncate(24 * time.Hour).Before(time.Now().Truncate(24 * time.Hour)) {
		*validator_errs = append(*validator_errs, ValidatorErr{
			Name: "DeadlineErr",
			Desc: "Deadline must not be in the past!",
		})
		return
	}
}

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

func CostValidator(cost float64, validator_errs *[]ValidatorErr) {
	if cost >= 0 {
		return
	} 

	*validator_errs = append(*validator_errs, ValidatorErr{
		Name: "CostErr",
		Desc: "Invalid cost value!",
	})
}

func HasResolverValidator(request_ref *uint, validator_errs *[]ValidatorErr) {
	if request_ref == nil {
		*validator_errs = append(*validator_errs, ValidatorErr{
			Name: "MaintenReqRefErr",
			Desc: "Reference to maintenance request is required!",
		})
		return
	}

	request_model := &models.MaintenanceRequest{}
	if result := utils.DB.First(request_model, *request_ref); result.Error != nil {
		*validator_errs = append(*validator_errs, ValidatorErr{
			Name: "MaintenReqRefErr",
			Desc: result.Error.Error(),
		})
		return
	}

	if request_model.ResolvedByRef == nil {
		*validator_errs = append(*validator_errs, ValidatorErr{
			Name: "MaintenReqRefErr",
			Desc: "Maintenance request has no resolver!",
		})
	}
}
