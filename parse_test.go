package harfileparser

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestReadingHarFileWithoutError(t *testing.T) {
	h := HarParser{}
	h.Init("./testdata/google.json")
	err := h.ParseHarFile()
	if err != nil {
		t.Errorf("Error reading file. %v", err)
	}
}

func TestParseAndGetEntriesCount(t *testing.T) {
	h := HarParser{}
	h.Init("./testdata/google.json")
	err := h.ParseHarFile()
	if err != nil {
		t.Errorf("Error reading file. %v", err)
	}
	ExpectedCount := 0
	if count := h.HarData.Log.EntriesCount(); count <= ExpectedCount {
		t.Errorf("Entires count should be Greater than %d.", ExpectedCount)
	}
}

func TestGetAllUrlsFunction(t *testing.T) {
	h := HarParser{}
	h.Init("./testdata/google.json")
	err := h.ParseHarFile()
	if err != nil {
		t.Errorf("Error reading file. %v", err)
	}
	ExpectedCount := 0
	urls := h.HarData.Log.GetRequestUrls()
	t.Logf("urls Retrieved:\n%s\n", strings.Join(urls, "\n"))
	if count := len(urls); count <= ExpectedCount {
		t.Errorf("Entires count should be Greater than %d.", ExpectedCount)
	}
}

func TestRedirectCount(t *testing.T) {
	h := HarParser{}
	h.Init("./testdata/google.json")
	err := h.ParseHarFile()
	if err != nil {
		t.Errorf("Error reading file. %v", err)
	}

	f, err := os.Open("./testdata/google.json")
	if err != nil {
		t.Errorf("Error reading file. %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	redirectCount := 0
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "status") && strings.Contains(scanner.Text(), "302") {
			redirectCount++
		}
	}
	t.Logf("Redirect Count : %d", redirectCount)
	if count := h.HarData.Log.GetRedirectCounts(); count < redirectCount {
		t.Errorf("Entires count should be Greater than %d.", redirectCount)
	}
}

func TestEntriesToTextFlow(t *testing.T) {
	h := HarParser{}
	h.Init("./testdata/google.json")
	err := h.ParseHarFile()
	if err != nil {
		t.Errorf("Error reading file. %v", err)
	}
	flow := h.HarData.Log.EntriesToFlowText()
	t.Logf("Text To Flow : \n%s", flow)

	urls := h.HarData.Log.GetRequestUrls()

	//Check if all the Urls are available in the content.
	for _, v := range urls {
		if !(strings.Contains(flow, v)) {
			t.Errorf("URL details missing. Expected %s in flow.", v)
		}
	}
}
