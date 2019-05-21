# parse-har
Library to parse HAR file.

Basic functionalities includes,
1. parsing a HAR file.
2. Perform simple operations on the HAR file.
3. Use this libary as an input for other algorithms.

> The main purpose of writing this library is to use the HAR output as one of the input for Security operations like URL classifications, Machine learning algorithms ... etc

## Usage 

```go

h := HarParser{} // initialize the struct
h.Init("./testdata/google.json") // initialize the file
err := h.ParseHarFile() // parse the file
if err != nil {
    t.Errorf("Error reading file. %v", err)
}

```

## Roadmap

Since the main purpose of building this library is for security purpose, i will adding more functionalities w.r.t parsing the har contents for getting more meaning insights.

e.g. Functions like EntriesToFlowText() , exports Har file to a network traffic flow. (Currently in Beta.) working on to improve this.

```go
	h := HarParser{}
	h.Init("./testdata/google.json")
	err := h.ParseHarFile()
	if err != nil {
		t.Errorf("Error reading file. %v", err)
	}
	flow := h.HarData.Log.EntriesToFlowText()
	t.Logf("Text To Flow : \n%s", flow)

	urls := h.HarData.Log.GetRequestUrls()
```