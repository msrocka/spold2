package spold2

import (
	"encoding/xml"
	"io/ioutil"
)

// ReadFile reads an EcoSpold 2 file.
func ReadFile(fileName string) (*EcoSpold, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	spold := &EcoSpold{}
	err = xml.Unmarshal(data, spold)
	return spold, err
}
