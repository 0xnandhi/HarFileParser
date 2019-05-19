package harfileparser

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//HarParser structure to hold Har parsed data.
type HarParser struct {
	FileName string
	HarData  *HAR
}

//Init Har parser initializing logic
func (h *HarParser) Init(filename string) {
	h.FileName = filename
}

func (h *HarParser) readConfigFile() []byte {
	content, err := ioutil.ReadFile(h.FileName)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

//ParseHarFile parses the json file and maps to sources struct
func (h *HarParser) ParseHarFile() error {
	jsonBlob := h.readConfigFile()
	sources := HAR{}
	err := json.Unmarshal(jsonBlob, &sources)
	h.HarData = &sources
	if err != nil {
		return err
	}
	return nil
}
