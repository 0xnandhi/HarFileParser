package harfileparser

import "testing"

func TestReadingHarFileWithoutError(t *testing.T) {
	h := HarParser{}
	h.Init("./testdata/google.json")
	err := h.ParseHarFile()
	if err != nil {
		t.Errorf("Error reading file. %v", err)
	}
}
