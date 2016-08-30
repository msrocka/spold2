package spold2

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
)

func TestReadModel(t *testing.T) {
	spold := readExample()
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
	spold := readExample()
	if spold.DataSet.FlowData == nil {
		t.Error("No flow data")
	}
	if len(spold.DataSet.FlowData.ElementaryExchanges) != 2 {
		t.Error("Missing elementary flows")
	}
}

func readExample() *EcoSpold {
	data, _ := ioutil.ReadFile("./testdata/example.xml")
	spold := &EcoSpold{}
	xml.Unmarshal(data, spold)
	return spold
}
