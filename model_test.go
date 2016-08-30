package spold2

import (
	"encoding/xml"
	"testing"
)

func TestReadModel(t *testing.T) {
	spold := example(t)
	if spold.DataSet.Description.ID != "0b647620-42e0-49d6-abd5-9cdde1b8fb01" {
		t.Error("Failed to read data set ID")
		return
	}
	data, err := xml.MarshalIndent(spold, "", "  ")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(data))
}

func TestElementaryExchanges(t *testing.T) {
	spold := example(t)
	if spold.DataSet.FlowData == nil {
		t.Error("No flow data")
	}
	if len(spold.DataSet.FlowData.ElementaryExchanges) != 2 {
		t.Error("Missing elementary flows")
	}
}

func example(t *testing.T) *EcoSpold {
	spold, err := ReadFile("./testdata/example.xml")
	if err != nil {
		t.Error(err)
	}
	return spold
}
