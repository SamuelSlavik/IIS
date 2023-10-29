package serializers

type ConnectionSerializer struct {
	ID        uint
	LineName  string
	Type      string
	ListStops *[]StopsSerializer
}

type StopsSerializer struct {
	DepartureTime string
	StopName      string
}
