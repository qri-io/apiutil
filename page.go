package apiutil

import "net/http"

// DefaultPageSize is the
const DefaultPageSize = 100

// Page represents pagination information
type Page struct {
	Number int `json:"page"`
	Size   int `json:"pageSize"`
}

// NewPage is a conveince page constructor
func NewPage(number, size int) Page {
	return Page{number, size}
}

// Limit is a convenience accessor for page size
func (p Page) Limit() int {
	return p.Size
}

// Offset calculates the starting index for pagination based on page
// size & number
func (p Page) Offset() int {
	return (p.Number - 1) * p.Size
}

// PageFromRequest extracts pagination params from an http request
func PageFromRequest(r *http.Request) Page {
	var number, size int
	if i, err := ReqParamInt("page", r); err == nil {
		number = i
	}
	if number <= 0 {
		number = 1
	}

	if i, err := ReqParamInt("pageSize", r); err == nil {
		size = i
	}
	if size <= 0 {
		size = DefaultPageSize
	}

	return NewPage(number, size)
}

// NewPageFromOffsetAndLimit converts a offset and Limit to a Page struct
func NewPageFromOffsetAndLimit(offset, limit int) Page {
	var number, size int
	size = limit
	if size <= 0 {
		size = DefaultPageSize
	}
	number = offset/size + 1
	return NewPage(number, size)
}
