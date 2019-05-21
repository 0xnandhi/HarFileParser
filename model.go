package harfileparser

import (
	"strconv"
	"strings"
)

// defines the har model structure as per har spec 1.2

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
	OnContentLoad float64 `json:"onContentLoad,omitempty"`
	OnLoad        float64 `json:"onLoad,omitempty"`
	Comment       string  `json:"comment,omitempty"`
}

//Page list of all exported pages.
type Page struct {
	StartedDateTime string     `json:"startedDateTime"`
	ID              string     `json:"id"`
	Tite            string     `json:"title"`
	PageTimings     PageTiming `json:"pageTimings"`
	Comment         string     `json:"comment,omitempty"`
}

//QueryString list of params passed from a query string.
type QueryString struct {
	Name    string `json:"name,omitempty"`
	Value   string `json:"value,omitempty"`
	Comment string `json:"comment,omitempty"`
}

//Params posted params in HTTP post request body
type Params struct {
	Name        string `json:"name"`
	Value       string `json:"value,omitempty"`
	FileName    string `json:"fileName,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	Comment     string `json:"comment,omitempty"`
}

//PostData describes the posted data in the HTTP request.
type PostData struct {
	MimeType string   `json:"mimeType"`
	Params   []string `json:"params"`
	Text     string   `json:"text,omitempty"`
	Comment  string   `json:"comment,omitempty"`
}

//Header req / response header information
type Header struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Comment string `json:"comment,omitempty"`
}

//Cookie cookie information structure
type Cookie struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Path     string `json:"path,omitempty"`
	Domain   string `json:"domain,omitempty"`
	Expires  string `json:"expires,omitempty"`
	HTTPOnly bool   `json:"httpOnly,omitempty"`
	Secure   bool   `json:"secure,omitempty"`
	Comment  string `json:"comment,omitempty"`
}

//Request detailed info about the HTTP request
type Request struct {
	Method      string        `json:"method"`
	URL         string        `json:"url"`
	HTTPVersion string        `json:"httpVersion"`
	Cookies     []Cookie      `json:"cookies"`
	Headers     []Header      `json:"headers"`
	QueryString []QueryString `json:"queryString"`
	PostData    PostData      `json:"postData,omitempty"`
	HeaderSize  int           `json:"headerSize"`
	BodySize    int           `json:"bodySize"`
	Comment     string        `json:"comment,omitempty"`
}

//ResponseContent describes the details about the responnse content text.
type ResponseContent struct {
	Size        int    `json:"size"`
	Compression int    `json:"compression,omitempty"`
	MimeType    string `json:"mimeType"`
	Text        string `json:"text,omitempty"`
	Encoding    string `json:"encoding,omitempty"`
	Comment     string `json:"comment,omitempty"`
}

//Response detailed info about the response.
type Response struct {
	Status       int             `json:"status"`
	StatusText   string          `json:"statusText"`
	HTTPVersion  string          `json:"httpVersion"`
	Cookies      []Cookie        `json:"cookies"`
	Headers      []Header        `json:"headers"`
	Content      ResponseContent `json:"content,omitempty"`
	RedirectURL  string          `json:"redirectURL"`
	HeaderSize   int             `json:"headerSize"`
	BodySize     int             `json:"bodySize"`
	TransferSize int             `json:"_transferSize,omitempty"`
	Comment      string          `json:"comment,omitempty"`
}

//Cache info about the cache usage
type Cache struct {
	BeforeRequest CacheRequest `json:"beforeRequest,omitempty"`
	AfterRequest  CacheRequest `json:"afterRequest,omitempty"`
	Comment       string       `json:"comment,omitempty"`
}

//CacheRequest Cache before / after request structure
type CacheRequest struct {
	Expires    string `json:"expires,omitempty"`
	LastAccess string `json:"lastAccess,omitempty"`
	ETag       string `json:"eTag,omitempty"`
	HitCount   int    `json:"hitCount,omitempty"`
	Comment    string `json:"comment,omitempty"`
}

//Timings detailed timing info about the request and response round trip
type Timings struct {
	Blocked int    `json:"blocked,omitempty"`
	DNS     int    `json:"dns,omitempty"`
	Connect int    `json:"connect,omitempty"`
	Send    int    `json:"send,omitempty"`
	Wait    int    `json:"wait,omitempty"`
	Receive int    `json:"receive,omitempty"`
	Comment string `json:"comment,omitempty"`
}

//Initiator details on which element invoked the request.
type Initiator struct {
}

//Entry list of all exported requests
type Entry struct {
	Pageref         string   `json:"pageref,omitempty"`
	StartedDateTime string   `json:"startedDateTime"`
	Time            float32  `json:"time"`
	Request         Request  `json:"request"`
	Response        Response `json:"response"`
	Cache           Cache    `json:"cache,omitempty"`
	FromDiskCache   bool     `json:"_fromDiskCache,omitempty"`
	Timings         Timings  `json:"timing"`
	ServerIPAddress string   `json:"serverIPAddress"`
	Connection      string   `json:"connection,omitempty"`
	Priority        string   `json:"_priority,omitempty"`
	Comment         string   `json:"comment,omitempty"`
}

//Log Root of the exported data
type Log struct {
	Version string  `json:"version"`
	Creator Creator `json:"creator"`
	Browser Browser `json:"browser,omitempty"`
	Pages   []Page  `json:"pages,omitempty"`
	Entries []Entry `json:"entries"`
	Comment string  `json:"comment,omitempty"`
}

//HAR structure as per spec 1.2
type HAR struct {
	Log Log `json:"log"`
}

//EntriesCount list total number of entires
func (l Log) EntriesCount() int {
	return len(l.Entries)
}

//GetRequestUrls collects all request URLs from the parsed file and return an
// array of urls.
func (l Log) GetRequestUrls() []string {
	urls := []string{}
	for _, v := range l.Entries {
		urls = append(urls, v.Request.URL)
	}
	return urls
}

//GetRedirectCounts returns an integer which displays the number of redirections
//through the request
func (l Log) GetRedirectCounts() int {
	rCount := 0
	for _, v := range l.Entries {
		if v.Response.Status == 302 {
			rCount++
		}
	}
	return rCount
}

// EntriesToFlowText (Experimental feature) Convert the Har to a text flow. similar to wireshark, tcp flow
// the structure resembles similar to tcp flow but not exactly replicates a network flow.
// The main purpose of this to quickly use the text as an input to other engines for further processing.
// e.g. could be URL classification engine.
func (l Log) EntriesToFlowText() string {
	var flow strings.Builder
	var reqURL strings.Builder
	var reqHeaders strings.Builder
	var rspHeaders strings.Builder

	// iterate over all the entries.
	for _, entry := range l.Entries {

		// Request URL structure
		// Method <space> URL <space> <HTTP Version>\r\n
		reqURL.WriteString(entry.Request.Method)
		reqURL.WriteString(" ")
		reqURL.WriteString(entry.Request.URL)
		reqURL.WriteString(" ")
		if strings.Contains(entry.Request.HTTPVersion, "h2") {
			reqURL.WriteString("HTTP/2")
			reqURL.WriteString("\r\n")
		}

		if strings.Contains(entry.Request.HTTPVersion, "http/") {
			version := strings.ToUpper(entry.Request.HTTPVersion)
			reqURL.WriteString(version)
			reqURL.WriteString("\r\n")
		}

		// Request Headers of type <Key>:<space><value>\r\n
		for _, header := range entry.Request.Headers {
			reqHeaders.WriteString(header.Name)
			reqHeaders.WriteString(": ")
			reqHeaders.WriteString(header.Value)
			reqHeaders.WriteString("\r\n")
		}

		//TODO: Post Data is not implemented. to be implemented

		// check if the method is of type post / any post data is available.
		reqHeaders.WriteString("\r\n")

		rspHeaders.WriteString(strconv.Itoa(entry.Response.Status))
		rspHeaders.WriteString(" ")
		if entry.Response.StatusText == "" {
			if entry.Response.Status == 200 {
				entry.Response.StatusText = "OK"
			}
		}
		rspHeaders.WriteString(entry.Response.StatusText)
		rspHeaders.WriteString("\r\n")

		// response Headers of type <Key>:<space><value>\r\n
		for _, rheader := range entry.Response.Headers {
			if rheader.Name == "status" {
				continue
			}
			rspHeaders.WriteString(rheader.Name)
			rspHeaders.WriteString(": ")
			rspHeaders.WriteString(rheader.Value)
			rspHeaders.WriteString("\r\n")
		}

		rspHeaders.WriteString("\r\n")

		// write all details to the flow
		flow.WriteString(reqURL.String())
		flow.WriteString(reqHeaders.String())
		flow.WriteString(rspHeaders.String())

		if len(entry.Response.Content.Text) > 0 {
			flow.WriteString(entry.Response.Content.Text)
		}
		flow.WriteString("\r\n")

		reqURL.Reset()
		reqHeaders.Reset()
		rspHeaders.Reset()

	}

	return flow.String()
}
