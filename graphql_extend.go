package graphql

import (
	"bufio"
	"strings"
	"bytes"
	"github.com/neelance/graphql-go/errors"
)

func ParseSchemaMultiFiles(contents map[string]string, resolver interface{}, opts ...SchemaOpt) (*Schema, error) {
	buffer := &bytes.Buffer{}
	sourceMap := errors.SourceMap{}
	nextStart := 0
	for filename, content := range contents {
		sourceMap.Add(errors.SourceInfo{Filename: filename, LineNumberStart: nextStart})
		content += "\n"
		line := countLine(content)
		nextStart += line
		buffer.WriteString(content)
	}

	s, err := ParseSchema(buffer.String(), resolver, opts...)
	if err != nil {
		if queryError, ok := err.(*errors.QueryError); ok {
			queryError.SourceMap = &sourceMap
		}
	}
	return s, err
}

func countLine(content string) int {
	s := bufio.NewScanner(strings.NewReader(content))
	count := 0
	for s.Scan() {
		count++
	}
	return count
}
