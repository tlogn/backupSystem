package filter

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Filter struct {
	Filename []string	`json:"filename"`
	Filetype []string	`json:"filetype"`
}

func (filter *Filter) ReadFilter(filePath string) error {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return err
	}
	err = json.Unmarshal(f, &filter)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (filter *Filter) IsLegal(filename string) bool {
	for _, name := range filter.Filename {
		if name == filename {
			return false
		}
	}
	for _, suffix := range filter.Filetype {
		if len(filename) >= len(suffix) && filename[len(filename) - len(suffix)  :] == suffix {
			return false
		}
	}
	return true
}