//nolint
package tankerkoenig

import (
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient("00000000-0000-0000-0000-000000000002", nil)
	if client == nil {
		t.Error("Client is nil")
	}
	if client.Station == nil {
		t.Error("StationService is nil")
	}
	if client.Prices == nil {
		t.Error("PricesService is nil")
	}
}

func TestStationList(t *testing.T) {
	client := NewClient("00000000-0000-0000-0000-000000000002", nil)
	stations, resp, err := client.Station.List(52.52099975265203, 13.43803882598877, 4)
	if err != nil {
		t.Errorf("returned an error. got=%s", err)
	}
	if resp == nil {
		t.Error("returned no response")
	}
	if len(stations) == 0 {
		t.Errorf("returned no stations. got=%d", len(stations))
	}
}

func TestStationDetail(t *testing.T) {
	client := NewClient("00000000-0000-0000-0000-000000000002", nil)
	id := "1c4f126b-1f3c-4b38-9692-05c400ea8e61"
	station, resp, err := client.Station.Detail(id)
	if err != nil {
		t.Errorf("returned an error. got=%s", err)
	}
	if resp == nil {
		t.Error("returned no response")
	}
	if reflect.DeepEqual(station, Station{}) {
		t.Error("returned empty station")
	}
	if station.Id != id {
		t.Errorf("returned the wrong station. got=%+v", station)
	}
}

func TestPricesGet(t *testing.T) {
	client := NewClient("00000000-0000-0000-0000-000000000002", nil)
	id := "1c4f126b-1f3c-4b38-9692-05c400ea8e61"
	prices, resp, err := client.Prices.Get(id)
	if err != nil {
		t.Errorf("returned an error. got=%s", err)
	}
	if resp == nil {
		t.Error("returned no response")
	}
	if len(prices) == 0 {
		t.Errorf("returned no prices. got=%d", len(prices))
	}
	if p, ok := prices[id]; !ok {
		t.Errorf("returned wrong prices. got=%+v", prices)
	} else {
		if reflect.DeepEqual(p, Price{}) {
			t.Error("returned empty price")
		}
	}
}
