package endpointdef

// TODO DONIA might be elsewhere ?
const (
	ContentType  = "Content-Type"
	MimeTypeJSON = "application/json"
	MimeTypeHTML = "text/html"
)

// Meta defines endpoint metadata
type Meta interface {
	TraceName() string
	Path() string
	Verb() string
	SuccessCode() int
}

// New creates a new endpoint structure
func New(traceName, path, verb string, code int) Meta {
	e := endpoint{
		traceName,
		path,
		verb,
		code,
	}
	return &e
}

// endpoint holds information about it
type endpoint struct {
	traceName   string
	path        string
	verb        string
	successCode int
}

// TraceName returns the trace name
func (e *endpoint) TraceName() string {
	return e.traceName
}

// Path returns the path
func (e *endpoint) Path() string {
	return e.path
}

// Verb returns the verb
func (e *endpoint) Verb() string {
	return e.verb
}

// SuccessCode returns the success code
func (e *endpoint) SuccessCode() int {
	return e.successCode
}
