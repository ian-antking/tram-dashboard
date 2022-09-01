package metrolink

type MetroLink struct{
	Id int
	Line string
	StationLocation string
	Direction string

	Dest0 string
	Carriages0 string
	Status0 string
	Wait0 string

	Dest1 string
	Carriages1 string
	Status1 string
	Wait1 string

	Dest2 string
	Carriages2 string
	Status2 string
	Wait2 string
}

type MetrolinkResponseBody struct {
	Value []MetroLink `json:"value"`
}