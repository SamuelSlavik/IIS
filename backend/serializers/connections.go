package serializers

import "time"

type ConnectionListSerializer struct {
	ArrivalTime   time.Time
	DepartureTime time.Time
	LineName      string
}
