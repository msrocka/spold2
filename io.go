package spold2

import (
	"encoding/xml"
	"io/ioutil"
	"path/filepath"
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

// EachFile parses each file in the given folder and calls the given handler
// with the data sets from respective files.
func EachFile(folder string, fn func(*ActivityDataSet) error) error {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return err
	}
	for _, info := range files {
		path := filepath.Join(folder, info.Name())
		spold, err := ReadFile(path)
		if err != nil {
			return err
		}
		if err = fn(spold.GetDataSet()); err != nil {
			return err
		}
	}
	return nil
}
