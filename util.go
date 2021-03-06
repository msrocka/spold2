package spold2

// RefFlow searches and returns the reference flow (product exchange with
// outputGroup=0) of a data set
func (s *EcoSpold) RefFlow() *IntermediateExchange {
	flows := s.IntermediateExchanges()
	if flows == nil {
		return nil
	}
	var candidate *IntermediateExchange
	for i := range flows {
		f := flows[i]
		if f.OutputGroup == nil {
			continue
		}
		group := *f.OutputGroup
		if group != 0 {
			continue
		}
		if f.Amount != 0 {
			return &f
		}
		if candidate == nil {
			candidate = &f
		}
	}
	return candidate
}

// IntermediateExchanges returns the product and waste flows of a data set
func (s *EcoSpold) IntermediateExchanges() []IntermediateExchange {
	dataSet := s.GetDataSet()
	if dataSet == nil {
		return nil
	}
	flows := dataSet.FlowData
	if flows == nil {
		return nil
	}
	return flows.IntermediateExchanges
}

// ElementaryExchanges returns the elementary flows of a data set
func (s *EcoSpold) ElementaryExchanges() []ElementaryExchange {
	dataSet := s.GetDataSet()
	if dataSet == nil {
		return nil
	}
	flows := dataSet.FlowData
	if flows == nil {
		return nil
	}
	return flows.ElementaryExchanges
}

// FindElementaryExchange searches for an elementary flow with the given
// name in the data set.
func (s *EcoSpold) FindElementaryExchange(name string) *ElementaryExchange {
	exchanges := s.ElementaryExchanges()
	for i := range exchanges {
		e := &exchanges[i]
		if e.Name == name {
			return e
		}
	}
	return nil
}

// FindProperty searches for a property with the given name in an exchange.
func (e *Exchange) FindProperty(name string) *Property {
	properties := e.Properties
	for i := range properties {
		p := &properties[i]
		if p.Name == name {
			return p
		}
	}
	return nil
}

// GetDataSet returns the data set or child data set that is stored in the
// EcoSpold document.
func (s *EcoSpold) GetDataSet() *ActivityDataSet {
	if s.DataSet == nil {
		return s.ChildDataSet
	}
	return s.DataSet
}
