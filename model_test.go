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

func TestExchangeGroup(t *testing.T) {
	spold := example(t)
	count := 0
	for _, e := range spold.ElementaryExchanges() {
		if e.Name == "1,1,1,2-Tetrachloroethane" {
			if *e.OutputGroup != 4 {
				t.Error("Wrong outputGroup for", e.Name)
			}
			if e.InputGroup != nil {
				t.Error("Wrong inputGroup for", e.Name)
			}
			count++
		}
		if e.Name == "Chlorine" {
			if *e.InputGroup != 4 {
				t.Error("Wrong inputGroup for", e.Name)
			}
			if e.OutputGroup != nil {
				t.Error("Wrong outputGroup for", e.Name)
			}
			count++
		}
	}
	if count != 2 {
		t.Error("Did not found all elementary flows")
	}
}

func TestCompartment(t *testing.T) {
	e := example(t).FindElementaryExchange("Chlorine")
	if e.Compartment.SubCompartment != "in water" {
		t.Error("Could not find compartment in", e.Name)
	}
}

func TestRefFlow(t *testing.T) {
	spold := example(t)
	ref := spold.RefFlow()
	if ref.Name != "Test product flow" {
		t.Error("Did not found reference flow")
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
