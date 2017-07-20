package errors

import (
	"fmt"
	"sort"
	"strings"
)

type SourceMap struct {
	sourceInfoList []SourceInfo
}

type SourceInfo struct {
	LineNumberStart int
	Filename        string
}

func (s *SourceMap) Add(info SourceInfo) {
	s.sourceInfoList = append(s.sourceInfoList, info)
	sort.Slice(s.sourceInfoList, func(i int, j int) bool {
		return s.sourceInfoList[i].LineNumberStart < s.sourceInfoList[j].LineNumberStart
	})
}

func (s *SourceMap) LineInFile(line int) (int, string) {
	var infoFound *SourceInfo
	for i, info := range s.sourceInfoList {
		if info.LineNumberStart > line {
			infoFound = &s.sourceInfoList[i-1]
		}
	}
	if infoFound == nil {
		infoFound = &s.sourceInfoList[len(s.sourceInfoList)-1]
	}
	return line - infoFound.LineNumberStart, infoFound.Filename
}

func (s *SourceMap) String() string {
	texts := []string{}
	for _, info := range s.sourceInfoList {
		texts = append(texts, fmt.Sprintf("%s:%d", info.Filename, info.LineNumberStart))
	}
	return fmt.Sprintf("[%s]", strings.Join(texts, ", "))
}
