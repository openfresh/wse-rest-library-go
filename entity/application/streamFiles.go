package application

import (
	"github.com/openfresh/wse-rest-library-go/entity/base"
)

type StreamFiles struct {
	base.EntityBase
	ID   string `json:"id"`
	Href string `json:"href"`
}

func NewStreamFiles() *StreamFiles {
	s := new(StreamFiles)
	return s
}

func (s *StreamFiles) EntityName() string {
	return "streamFiles"
}

func (s *StreamFiles) SetURI(baseURI string) {
	s.SetRestURI("")
}
