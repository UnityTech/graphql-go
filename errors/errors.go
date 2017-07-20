package errors

import (
	"fmt"
)

type QueryError struct {
	Message       string        `json:"message"`
	Locations     []Location    `json:"locations,omitempty"`
	Path          []interface{} `json:"path,omitempty"`
	Rule          string        `json:"-"`
	ResolverError error         `json:"-"`
	SourceMap     *SourceMap    `json:"-"`
}

type Location struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

func (a Location) Before(b Location) bool {
	return a.Line < b.Line || (a.Line == b.Line && a.Column < b.Column)
}

func Errorf(format string, a ...interface{}) *QueryError {
	return &QueryError{
		Message: fmt.Sprintf(format, a...),
	}
}

func (err *QueryError) Error() string {
	if err == nil {
		return "<nil>"
	}
	str := fmt.Sprintf("graphql: %s", err.Message)
	for _, loc := range err.Locations {
		if err.SourceMap == nil {
			str += fmt.Sprintf(" (line %d, column %d)", loc.Line, loc.Column)
		} else {
			line, file := err.SourceMap.LineInFile(loc.Line)
			str += fmt.Sprintf(" (file %s, line %d, column %d)", file, line, loc.Column)
		}

	}
	return str
}

var _ error = &QueryError{}
