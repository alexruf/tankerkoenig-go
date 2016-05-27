package tankerkoenig

import (
	"fmt"
	"net/url"
)

// StationService is an interface to query station information from the Tankerkönig-API.
type StationService interface {
	List(lat float64, lng float64, rad int) ([]Station, *Response, error)
}

// StationServiceOp handles communication with the station related methods of the Tankerkönig-API.
type StationServiceOp struct {
	client *Client
}

var _ StationService = &StationServiceOp{}

// Station represents a gas station.
type Station struct {
	Brand       string      `json:"brand"`       // Brand
	Dist        float64     `json:"dist"`        // Distance (air line) from the search point to the gas station
	HouseNumber string      `json:"houseNumber"` // House number
	Id          string      `json:"id"`          // ID
	IsOpen      bool        `json:"isOpen"`      // Open-status
	Lat         float64     `json:"lat"`         // Latitude
	Lng         float64     `json:"Lng"`         // Longitude
	Name        string      `json:"name"`        // Name
	Place       string      `json:"place"`       // Place
	PostCode    int         `json:"postCode"`    // Post code
	Diesel      interface{} `json:"diesel"`      // Price for diesel fuel type
	E5          interface{} `json:"e5"`          // Price for E5 fuel type
	E10         interface{} `json:"e10"`         // Price for E10 fuel type
	Street      string      `json:"street"`      // Street
}

// stationRoot represents a response from the Tankerkönig-API.
type stationsRoot struct {
	Status   string    `json:"status"`
	Ok       bool      `json:"ok"`
	License  string    `json:"license"`
	Data     string    `json:"data"`
	Stations []Station `json:"stations"`
}

// List returns all stations within a radius of a location.
func (s *StationServiceOp) List(lat float64, lng float64, rad int) ([]Station, *Response, error) {
	path := "json/list.php"

	query := url.Values{}
	query.Add("lat", fmt.Sprintf("%.13f", lat))
	query.Add("lng", fmt.Sprintf("%.13f", lng))
	query.Add("rad", fmt.Sprintf("%d", rad))
	query.Add("type", "all")
	query.Add("apikey", s.client.APIKey)
	query.Add("sort", "dist")

	req, err := s.client.NewRequest("GET", path, query, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(stationsRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, nil, err
	}

	return root.Stations, resp, nil
}
