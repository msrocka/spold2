package masterdata

import (
	"encoding/xml"
	"io/ioutil"
)

type CompartmentList struct {
	XMLName      xml.Name      `xml:"http://www.EcoInvent.org/EcoSpold02 validCompartments"`
	Compartments []Compartment `xml:"compartment"`
}

type Compartment struct {
	ID              string           `xml:"id,attr"`
	Name            string           `xml:"name"`
	SubCompartments []SubCompartment `xml:"subcompartment"`
}

type SubCompartment struct {
	ID      string `xml:"id,attr"`
	Name    string `xml:"name"`
	Comment string `xml:"comment"`
}

func ReadCompartments(file string) (*CompartmentList, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	list := &CompartmentList{}
	if err := xml.Unmarshal(data, list); err != nil {
		return nil, err
	}
	return list, nil
}
