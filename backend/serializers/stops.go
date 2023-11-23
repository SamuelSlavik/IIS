package serializers

import "github.com/AdamPekny/IIS/backend/models"

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

func (s *StopSerializer) FromModel(stop models.Stop) {
	s.ID = stop.ID
	s.Name = stop.Name
	s.Active = true

	// TODO - implement active status
}
