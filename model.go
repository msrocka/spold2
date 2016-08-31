package spold2

import (
	"encoding/xml"
)

// EcoSpold is the container for activity data sets.
type EcoSpold struct {
	XMLName      xml.Name         `xml:"http://www.EcoInvent.org/EcoSpold02 ecoSpold"`
	DataSet      *ActivityDataSet `xml:"activityDataset,omitempty"`
	ChildDataSet *ActivityDataSet `xml:"childActivityDataset,omitempty"`
}

// ActivityDataSet contains the information of an activity.
type ActivityDataSet struct {
	Description     *ActivityDescription `xml:"activityDescription>activity"`
	Classifications []Classification     `xml:"activityDescription>classification"`
	Geography       *Geography           `xml:"activityDescription>geography"`
	Technology      *Technology          `xml:"activityDescription>technology"`
	TimePeriod      *TimePeriod          `xml:"activityDescription>timePeriod"`
	FlowData        *FlowData            `xml:"flowData,omitempty"`
}

// ActivityDescription contains the identifying information of an activity
// dataset including name and classification.
type ActivityDescription struct {
	ID               string `xml:"id,attr"`
	NameID           string `xml:"activityNameId,attr,omitempty"`
	ParentID         string `xml:"parentActivityId,attr,omitempty"`
	InheritanceDepth int    `xml:"inheritanceDepth,attr"`
	Type             int    `xml:"type,attr"`
	SpecialType      int    `xml:"specialActivityType,attr"`

	Name    string        `xml:"activityName"`
	Comment *TextAndImage `xml:"generalComment,omitempty"`
}

// Classification contains classification pairs to specify the activity.
type Classification struct {
	ID string `xml:"classificationId,attr"`

	System string `xml:"classificationSystem,omitempty"`
	Value  string `xml:"classificationValue,omitempty"`
}

// Geography contains information about the geographic validity of the process.
type Geography struct {
	ID string `xml:"geographyId,attr"`

	ShortName string        `xml:"shortname,omitempty"`
	Comment   *TextAndImage `xml:"comment,omitempty"`
}

// Technology contains information about the process technology.
type Technology struct {
	Level   int           `xml:"technologyLevel,attr"`
	Comment *TextAndImage `xml:"comment,omitempty"`
}

// TimePeriod contains information about the process time.
type TimePeriod struct {
	StartDate                  string        `xml:"startDate,attr"`
	EndDate                    string        `xml:"endDate,attr"`
	IsDataValidForEntirePeriod bool          `xml:"isDataValidForEntirePeriod,attr"`
	Comment                    *TextAndImage `xml:"comment,omitempty"`
}

// TODO: time etc.

// FlowData contains the inputs and outputs of an activity data set.
type FlowData struct {
	IntermediateExchanges []IntermediateExchange `xml:"intermediateExchange"`
	ElementaryExchanges   []ElementaryExchange   `xml:"elementaryExchange"`
}

// Exchange contains the common information of intermediate and elementary
// exchanges.
type Exchange struct {
	ID         string     `xml:"id,attr"`
	UnitID     string     `xml:"unitId,attr"`
	Amount     float64    `xml:"amount,attr"`
	Name       string     `xml:"name"`
	UnitName   string     `xml:"unitName"`
	Comment    string     `xml:"comment,omitempty"`
	Properties []Property `xml:"property"`
}

// IntermediateExchange comprises intermediate product and waste inputs and
// outputs for the activity.
type IntermediateExchange struct {
	Exchange
	ExchangeID     string `xml:"intermediateExchangeId,attr"`
	ActivityLinkID string `xml:"activityLinkId,attr"`
}

// ElementaryExchange are environmental inputs and outputs of a process
type ElementaryExchange struct {
	Exchange
	CAS        string `xml:"casNumber,attr,omitempty"`
	ExchangeID string `xml:"elementaryExchangeId,attr"`
}

// Property contains additional information about an exchange.
type Property struct {
	ID       string  `xml:"propertyId,attr"`
	Amount   float64 `xml:"amount,attr"`
	UnitID   string  `xml:"unitId,attr"`
	Name     string  `xml:"name"`
	UnitName string  `xml:"unitName"`
	Comment  string  `xml:"comment,omitempty"`
}
