package harfileparser

// defines the har model structure as per har spec 1.2

//HAR structure as per spec 1.2
type HAR struct {
}

//Creator Name and application of the log creator application.
type Creator struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	Comment string `json:"comment,omitempty"`
}

//Browser Name and version info of the browser used. Optional
type Browser struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Comment string `json:"comment,omitempty"`
}

//PageTiming detailed timing info about the page load.
type PageTiming struct {
	OnContentLoad int    `json:"onContentLoad,omitempty"`
	OnLoad        int    `json:"onLoad,omitempty"`
	Comment       string `json:"comment,omitempty"`
}

//Pages list of all exported pages.
type Pages struct {
	StartedDateTime string     `json:"startedDateTime"`
	ID              string     `json:"id"`
	Tite            string     `json:"title"`
	PageTimings     PageTiming `json:"pageTimings"`
	Comment         string     `json:"comment,omitempty"`
}

//Request detailed info about the request
type Request struct {
}

//Response detailed info about the response.
type Response struct {
}

//Entries list of all exported requests
type Entries struct {
	Pageref         string `json:"pageref,omitempty"`
	StartedDateTime string `json:"startedDateTime"`
	Time            int    `json:"time"`
}

//Comment provided by user or the application
type Comment struct {
}

//Log Root of the exported data
type Log struct {
	Version string  `json:"version"`
	Creator Creator `json:"creator"`
	Browser Browser `json:"browser,omitempty"`
	Pages   Pages   `json:"pages,omitempty"`
	Entries Entries `json:"entires"`
	Comment Comment `json:"comment,omitempty"`
}
