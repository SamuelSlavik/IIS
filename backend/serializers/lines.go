package serializers

import (
	"github.com/AdamPekny/IIS/backend/models"
	"github.com/AdamPekny/IIS/backend/utils"
)

type LineInList struct {
	Name        string
	InitialStop string
	FinalStop   string
}

type LineSerializer struct {
	Name          string
	StopsSequence []string
}

type LineCreateSerializer struct {
	Name          string `binding:"required"`
	StopsSequence []CreateSeqStops
}
type CreateSeqStops struct {
	StopName string `binding:"required"`
	Duration uint   `binding:"required"`
}

func (line_s *LineSerializer) GetStops(line_name string) error {
	var line models.Line
	if err := utils.DB.Model(&line).Preload("Segments").First(&line, "Name = ?", line_name).Error; err != nil {
		return err
	}
	for i := 0; i < len(line.Segments); i++ {
		if line.Segments[i].StopName2 == line.FinalStop {
			line_s.StopsSequence = append(line_s.StopsSequence, line.Segments[i].StopName1)
			line_s.StopsSequence = append(line_s.StopsSequence, line.Segments[i].StopName2)
			break
		}
		line_s.StopsSequence = append(line_s.StopsSequence, line.Segments[i].StopName1)

	}
	return nil
}

func (l *LineInList) FromModel(line models.Line) {
	l.Name = line.Name
	l.InitialStop = line.InitialStop
	l.FinalStop = line.FinalStop
}
