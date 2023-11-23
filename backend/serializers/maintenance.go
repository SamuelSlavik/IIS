package serializers

import (
	"github.com/AdamPekny/IIS/backend/models"
)



type MalfuncRepCreateSerialzier struct {
	Description string `binding:"required"`
	Status models.Status
	CreatedByRef uint
	VehicleRef uint `binding:"required"`
}

func (m *MalfuncRepCreateSerialzier) ToModel() *models.MalfunctionReport {
	malfunc_report := &models.MalfunctionReport{
		Description: m.Description,
		Status: m.Status,	
		CreatedByRef: m.CreatedByRef,
		VehicleRef: m.VehicleRef,
	}

	return malfunc_report
}

type MalfuncRepPublicSerialzier struct {
	Description string `binding:"required"`
	Status models.Status `binding:"required"`
	CreatedBy UserPublicSerializer `binding:"required"`
	Vehicle VehicleSerializer `binding:"required"`
}

func (m *MalfuncRepPublicSerialzier) FromModel(malfunc_report models.MalfunctionReport) (err error) {
	m.Description = malfunc_report.Description
	m.Status = malfunc_report.Status

	var created_by_public UserPublicSerializer

	if err := created_by_public.FromModel(malfunc_report.CreatedBy); err != nil {
		return err
	}

	m.CreatedBy = created_by_public

	// TODO Serialize vehicle

	return nil
} 


