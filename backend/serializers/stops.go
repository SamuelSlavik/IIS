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

func (u *StopSerializer) FromModel(stop models.Stop) {
	u.ID = stop.ID
	u.Name = stop.Name
	u.Active = true

	// TODO - implement active status
}
