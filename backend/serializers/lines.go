package serializers

import "github.com/AdamPekny/IIS/backend/models"

type LineInList struct {
	Name        string
	InitialStop string
	FinalStop   string
}

func (l *LineInList) FromModel(line models.Line) {
	l.Name = line.Name
	l.InitialStop = line.InitialStop
	l.FinalStop = line.FinalStop
}
