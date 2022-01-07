package masterdata

import "testing"

func TestReadCompartments(t *testing.T) {
	list, err := ReadCompartments("../testdata/compartments.xml")
	if err != nil {
		t.Error(err)
	}

	var air *Compartment
	for i := range list.Compartments {
		comp := list.Compartments[i]
		if comp.Name == "air" {
			air = &comp
			break
		}
	}
	if air == nil {
		t.Error("could not find compartment 'air'")
		return
	}

	var indoor *SubCompartment
	for i := range air.SubCompartments {
		sub := air.SubCompartments[i]
		if sub.Name == "indoor" {
			indoor = &sub
			break
		}
	}
	if indoor == nil {
		t.Error("could not find sub-compartment 'indoor'")
	}

}
