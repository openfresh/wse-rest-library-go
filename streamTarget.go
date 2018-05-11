package wserest

import (
	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
	"github.com/openfresh/wse-rest-library-go/entity/base"
)

// StreamTarget is PushPublish map entries utility
type StreamTarget struct {
	wowza
}

// NewStreamTarget create StreamTarget object
func NewStreamTarget(settings *helper.Settings, appName string) *StreamTarget {
	s := new(StreamTarget)
	s.init(settings)
	s.props["appName"] = appName
	s.baseURI = s.host() + "/servers/" + s.serverInstance() + "/vhosts/" + s.vHostInstance() + "/applications/" + appName + "/pushpublish/mapentries"

	return s
}

// Create adds the specified PushPublish map entry for the specified Application
func (s *StreamTarget) Create(
	sourceStreamName,
	entryName,
	profile,
	host,
	userName,
	password,
	streamName,
	application string) (map[string]interface{}, error) {
	s.setRestURI(s.baseURI + "/" + entryName)
	if sourceStreamName != "" {
		s.props["sourceStreamName"] = sourceStreamName
	}
	if entryName != "" {
		s.props["entryName"] = entryName
	}
	if profile != "" {
		s.props["profile"] = profile
	}
	if host != "" {
		s.props["host"] = host
	}
	if userName != "" {
		s.props["userName"] = userName
	}
	if password != "" {
		s.props["password"] = password
	}
	if streamName != "" {
		s.props["streamName"] = streamName
	}
	if application != "" {
		s.props["application"] = application
	}

	response, err := s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, POST, "")

	return response, err
}

// GetAll retrieves the list of PushPublish map entries for the specified Application
func (s *StreamTarget) GetAll() (map[string]interface{}, error) {
	s.setNoParams()
	s.setRestURI(s.baseURI)
	return s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

func (s *StreamTarget) setNoParams() {
	s.AddSkipParameter("userName") //todo: correct key name?
	s.AddSkipParameter("password")
	s.AddSkipParameter("group")
	s.AddSkipParameter("sourceStreamName")
	s.AddSkipParameter("entryName")
	s.AddSkipParameter("profile")
	s.AddSkipParameter("host")
	s.AddSkipParameter("application")
	s.AddSkipParameter("streamName")
}

// Remove deletes the specified PushPublish map entry for the specified Application
func (s *StreamTarget) Remove(entryName string) (map[string]interface{}, error) {
	s.setNoParams()
	s.setRestURI(s.baseURI + "/" + entryName)

	return s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, DELETE, "")
}
