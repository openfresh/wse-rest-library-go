package application

import (
	"github.com/openfresh/wse-rest-library-go/entity/base"
)

type StreamConfig struct {
	base.EntityBase
	StreamType           string   `json:"streamType"`
	LiveStreamPacketizer []string `json:"liveStreamPacketizer"`
}

func NewStreamConfig() *StreamConfig {
	s := new(StreamConfig)
	s.StreamType = "live"
	s.LiveStreamPacketizer = []string{"cupertinostreamingpacketizer", "smoothstreamingpacketizer", "sanjosestreamingpacketizer"}
	return s
}

func (s *StreamConfig) EntityName() string {
	return "streamConfig"
}

func (s *StreamConfig) SetURI(baseURI string) {
	s.SetRestURI(baseURI + "/streamconfiguration")
}
