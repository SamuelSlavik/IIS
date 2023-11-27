// package serializers holds structures and functions for serializing data
// this file contains serializers for stops
package serializers

import (
	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
)

// StopSerializer is used to serialize data about stop
// it is used in GET request to get data about stop
type StopSerializer struct {
	ID     uint
	Name   string
	Active bool
}

// EditConnection is used to serialize data about stop in connection
// it is used in PUT request to get data about stop in connection
type EditStopSerializer struct {
	ID   uint
	Name string
}

// StopCreateSerializer is used to serialize data about stop
// it is used in POST request to create a new stop
type StopCreateRequest struct {
	Name string `binding:"required"`
}

// FromModel is used to serialize data about stop
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
