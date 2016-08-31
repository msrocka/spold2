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

func TestTechnologyAndTime(t *testing.T) {
	spold := example(t)
	techComment := spold.DataSet.Technology.Comment.Texts[0].Value
	if techComment != "tech. desc." {
		t.Error("Technology comment not found")
		return
	}
	time := spold.DataSet.TimePeriod
	if time.StartDate != "2016-03-04" || time.EndDate != "2016-03-25" {
		t.Error("Could not read time period")
		return
	}
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
