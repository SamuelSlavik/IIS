package serializers

import "time"

type ConnectionListSerializer struct {
	ID            uint
	ArrivalTime   time.Time
	DepartureTime time.Time
	LineName      string
}

type ShowConnectionSerializer struct {
	DepartureTime string
	StopName      string
}
