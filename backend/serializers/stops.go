package serializers

import (
	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
)

type StopSerializer struct {
	ID     uint
	Name   string
	Active bool
}

type EditStopSerializer struct {
	ID   uint
	Name string
}

type StopCreateRequest struct {
	Name string `binding:"required"`
}

func (s *StopSerializer) FromModel(stop models.Stop) error {
	s.ID = stop.ID
	s.Name = stop.Name
	var segments []models.Segment
	result := utils.DB.Where("stop_name1 = ? OR stop_name2 = ?", stop.Name, stop.Name).Find(&segments)
	if result.Error != nil {
		return result.Error
	}
	if len(segments) > 0 {
		s.Active = true
	} else {
		s.Active = false
	}
	return nil
}
