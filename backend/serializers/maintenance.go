package serializers

import (
	"fmt"
	"time"

	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
	"github.com/AdamPekny/IIS/backend/validators"
	"github.com/gin-gonic/gin"
)



type MalfuncRepCreateSerialzier struct {
	Title string `binding:"required"`
	Description string `binding:"required"`
	CreatedByRef *uint
	VehicleRef *string `binding:"required"`
	ValidatorErrs []validators.ValidatorErr
}

func (m *MalfuncRepCreateSerialzier) ToModel() *models.MalfunctionReport {
	malfunc_report := &models.MalfunctionReport{
		Title: m.Title,
		Description: m.Description,
		CreatedByRef: m.CreatedByRef,
		VehicleRef: m.VehicleRef,
	}

	return malfunc_report
}

type MalfuncRepPublicSerialzier struct {
	ID uint `binding:"required"`
	Title string `binding:"required"`
	Description string `binding:"required"`
	CreatedBy *UserMaintenanceSerializer `binding:"required"`
	Vehicle *VehicleMaintenanceSerializer `binding:"required"`
	CreatedAt time.Time
}

func (m *MalfuncRepPublicSerialzier) FromModel(malfunc_report *models.MalfunctionReport) (err error) {
	m.ID = malfunc_report.ID
	m.Title = malfunc_report.Title
	m.Description = malfunc_report.Description
	m.CreatedAt = malfunc_report.CreatedAt

	created_by_serializer := &UserMaintenanceSerializer{}

	if err := created_by_serializer.FromModel(malfunc_report.CreatedBy); err != nil {
		return err
	}

	m.CreatedBy = created_by_serializer

	vehicle_serializer := &VehicleMaintenanceSerializer{}

	if err := vehicle_serializer.FromModel(malfunc_report.Vehicle); err != nil {
		return err
	}

	m.Vehicle = vehicle_serializer

	return nil
} 


type MalfuncRepShortPublicSerialzier struct {
	ID uint `binding:"required"`
	Title string `binding:"required"`
	CreatedByEmail string `binding:"required"`
	VehicleRef *string `binding:"required"`
	CreatedAt time.Time
}

func (m *MalfuncRepShortPublicSerialzier) FromModel(malfunc_report *models.MalfunctionReport) (err error) {
	m.ID = malfunc_report.ID
	m.Title = malfunc_report.Title
	m.VehicleRef = malfunc_report.VehicleRef
	m.CreatedAt = malfunc_report.CreatedAt


	if malfunc_report.CreatedBy == nil {
		malfunc_report.CreatedBy = &models.User{} 
		if result := utils.DB.First(malfunc_report.CreatedBy); result.Error != nil {
			return result.Error
		}
	}

	m.CreatedByEmail = malfunc_report.CreatedBy.Email

	return nil
}


type MaintenReqCreateSerializer struct {
	Status models.Status
	Deadline utils.CustomDate
	MalfuncRepRef *uint `binding:"required"`
	CreatedByRef *uint
	ResolvedByRef *uint
	ValidatorErrs []validators.ValidatorErr
}

func (m *MaintenReqCreateSerializer) Valid() bool {
	validators.StatusValidator(string(m.Status), &m.ValidatorErrs)
	if m.ResolvedByRef != nil {
		validators.HasRoleValidator(*m.ResolvedByRef, &m.ValidatorErrs, models.TechnicianRole)
	}

	return len(m.ValidatorErrs) == 0
}

func (m *MaintenReqCreateSerializer) ToModel(ctx *gin.Context) (*models.MaintenanceRequest, error) {
	model := &models.MaintenanceRequest{
		Status: m.Status,
		Deadline: m.Deadline.Time,
		MalfuncRepRef: m.MalfuncRepRef,
		ResolvedByRef: m.ResolvedByRef,
	}

	user, err := models.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	model.CreatedByRef = &user.ID

	return model, nil
}

type MaintenReqPublicSerializer struct {
	ID uint `binding:"required"`
	Status models.Status `binding:"required"`
	Deadline time.Time `binding:"required"`
	CreatedAt time.Time `binding:"required"`
	MalfuncRep *MalfuncRepShortPublicSerialzier `binding:"required"`
	CreatedBy *UserMaintenanceSerializer `binding:"required"`
	ResolvedBy *UserMaintenanceSerializer `binding:"required"`
	ValidatorErrs []validators.ValidatorErr
}

func (m *MaintenReqPublicSerializer) Valid() bool {
	validators.StatusValidator(string(m.Status), &m.ValidatorErrs)

	return len(m.ValidatorErrs) == 0
}

func (m *MaintenReqPublicSerializer) FromModel(mainten_req_model *models.MaintenanceRequest) (err error) {
	m.ID = mainten_req_model.ID
	m.Status = mainten_req_model.Status
	m.Deadline = mainten_req_model.Deadline
	m.CreatedAt = mainten_req_model.CreatedAt
	
	fmt.Printf("MalfuncRep ID: %d\n", mainten_req_model.MalfuncRep.ID)
	malfunc_rep_serializer := &MalfuncRepShortPublicSerialzier{}

	if err := malfunc_rep_serializer.FromModel(mainten_req_model.MalfuncRep); err != nil {
		return err
	}

	m.MalfuncRep = malfunc_rep_serializer

	created_by_serializer := &UserMaintenanceSerializer{}

	if err := created_by_serializer.FromModel(mainten_req_model.CreatedBy); err != nil {
		return err
	}

	m.CreatedBy = created_by_serializer

	if mainten_req_model.ResolvedByRef != nil {
		resolved_by_serializer := &UserMaintenanceSerializer{}

		if err := resolved_by_serializer.FromModel(mainten_req_model.ResolvedBy); err != nil {
			return err
		}

		m.ResolvedBy = resolved_by_serializer
	}
	

	return nil
}

type MaintenReqUpdateSerializer struct {
	Status models.Status `binding:"required"`
	Deadline utils.CustomDate `binding:"required"`
	MalfuncRepRef *uint `binding:"required"`
	CreatedByRef *uint
	ResolvedByRef *uint
	ValidatorErrs []validators.ValidatorErr
}

func (m *MaintenReqUpdateSerializer) Valid() bool {
	validators.StatusValidator(string(m.Status), &m.ValidatorErrs)
	if m.ResolvedByRef != nil {
		validators.HasRoleValidator(*m.ResolvedByRef, &m.ValidatorErrs, models.TechnicianRole)
	}

	return len(m.ValidatorErrs) == 0
}

func (m *MaintenReqUpdateSerializer) ToModel(ctx *gin.Context) (*models.MaintenanceRequest, error) {
	id, err := utils.GetIDFromURL(ctx)
	if err != nil {
		return nil, err
	}

	model := &models.MaintenanceRequest{}
	if result := utils.DB.First(model, id); result.Error != nil {
		return nil, result.Error
	}

	model.Status = m.Status
	model.Deadline = m.Deadline.Time
	model.MalfuncRepRef = m.MalfuncRepRef
	model.ResolvedByRef = m.ResolvedByRef

	return model, nil
}
