package metrolink

type MetroLink struct {
	Id              int    `json:"Id"`
	Line            string `json:"Line"`
	TLAREF          string `json:"TLAREF"`
	PIDREF          string `json:"PIDREF"`
	AtcoCode        string `json:"AtcoCode"`
	StationLocation string `json:"StationLocation"`
	Direction       string `json:"Direction"`

	Dest0      string `json:"Dest0"`
	Carriages0 string `json:"Carriages0"`
	Status0    string `json:"Status0"`
	Wait0      string `json:"Wait0"`

	Dest1      string `json:"Dest1"`
	Carriages1 string `json:"Carriages1"`
	Status1    string `json:"Status1"`
	Wait1      string `json:"Wait1"`

	Dest2      string `json:"Dest2"`
	Carriages2 string `json:"Carriages2"`
	Status2    string `json:"Status2"`
	Wait2      string `json:"Wait2"`

	Dest3      string `json:"Dest3"`
	Carriages3 string `json:"Carriages3"`
	Status3    string `json:"Status3"`
	Wait3      string `json:"Wait3"`

	MessageBoard string `json:"MessageBoard"`
	LastUpdated  string `json:"LastUpdated"`
}

type ResponseBody struct {
	DataContext string      `json:"@odata.context"`
	Value       []MetroLink `json:"value"`
}

type Tram struct {
	Destination string `json:"destination"`
	Carriages   string `json:"carriages"`
	Status      string `json:"status"`
	Wait        string `json:"wait"`
	LastUpdated string `json:"last_updated"`
}

type TramsResponseBody struct {
	Trams    []Tram   `json:"trams"`
	Messages []string `json:"messages"`
}
