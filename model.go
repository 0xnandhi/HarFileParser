package harfileparser

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
	Text        []byte `json:"text,omitempty"`
	Emcoding    string `json:"encoding,omitempty"`
	Comment     string `json:"comment,omitempty"`
}

//Response detailed info about the response.
type Response struct {
	Status      int             `json:"status"`
	StatusText  string          `json:"statusText"`
	HTTPVersion string          `json:"httpVersion"`
	Cookies     []Cookie        `json:"cookies"`
	Headers     []Header        `json:"headers"`
	Content     ResponseContent `json:"content,omitempty"`
	RedirectURL string          `json:"redirectURL"`
	HeaderSize  int             `json:"headerSize"`
	BodySize    int             `json:"bodySize"`
	Comment     string          `json:"comment,omitempty"`
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

//Entries list of all exported requests
type Entries struct {
	Pageref         string   `json:"pageref,omitempty"`
	StartedDateTime string   `json:"startedDateTime"`
	Time            int      `json:"time"`
	Request         Request  `json:"request"`
	Response        Response `json:"response"`
	Cache           Cache    `json:"cache,omitempty"`
	Timings         Timings  `json:"timing"`
	ServerIPAddress string   `json:"serverIPAddress"`
	Connection      string   `json:"connection,omitempty"`
	Comment         string   `json:"comment,omitempty"`
}

//Log Root of the exported data
type Log struct {
	Version string  `json:"version"`
	Creator Creator `json:"creator"`
	Browser Browser `json:"browser,omitempty"`
	Pages   Pages   `json:"pages,omitempty"`
	Entries Entries `json:"entires"`
	Comment string  `json:"comment,omitempty"`
}

//HAR structure as per spec 1.2
type HAR struct {
	Log Log `json:"log"`
}
