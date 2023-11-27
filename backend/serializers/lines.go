// package serializers holds structures and functions for serializing data
// this file contains serializers for lines
package serializers

import (
	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
)

// LineInList is used to serialize data about line
// it is used in GET request to get data about line
type LineInList struct {
	Name        string
	InitialStop string
	FinalStop   string
}

// LineSerializer is used to serialize data about line
// it is used in GET request to get data about line
type LineSerializer struct {
	Name          string
	StopsSequence []CreateSeqStops
}

// LineCreateSerializer is used to serialize data about line
// it is used in POST request to create a new line
type LineCreateSerializer struct {
	Name          string `binding:"required"`
	StopsSequence []CreateSeqStops
}

// CreateSeqStops is used to serialize stops in line
// it is used in POST request to create line segments
type CreateSeqStops struct {
	StopName string `binding:"required"`
	Duration uint   `binding:"required"`
}

// LineUpdateSerializer is used to serialize stops for line update
// it is used in PATCH request to update line segments
type LineUpdateSerializer struct {
	StopsSequence []CreateSeqStops
}

// GetStops gets serializes stops for line
func (line_s *LineSerializer) GetStops(line_name string) error {
	var line models.Line
	if err := utils.DB.Model(&line).Preload("Segments").First(&line, "Name = ?", line_name).Error; err != nil {
		return err
	}
	for i := 0; i < len(line.Segments); i++ {
		if line.Segments[i].StopName2 == line.FinalStop {
			line_s.StopsSequence = append(line_s.StopsSequence, CreateSeqStops{
				StopName: line.Segments[i].StopName1,
				Duration: line.Segments[i].Time,
			})
			line_s.StopsSequence = append(line_s.StopsSequence, CreateSeqStops{
				StopName: line.Segments[i].StopName2,
				Duration: 0,
			})
			break
		}
		line_s.StopsSequence = append(line_s.StopsSequence, CreateSeqStops{
			StopName: line.Segments[i].StopName1,
			Duration: line.Segments[i].Time,
		})

	}
	return nil
}

// FromModel loads data from model into serializer
func (l *LineInList) FromModel(line models.Line) {
	l.Name = line.Name
	l.InitialStop = line.InitialStop
	l.FinalStop = line.FinalStop
}
