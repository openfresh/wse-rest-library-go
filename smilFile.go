package wserest

import (
	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
	"github.com/openfresh/wse-rest-library-go/entity/base"
)

// SmilFile is SMIL Files utitility
type SmilFile struct {
	wowza
}

// NewSmilFile create SmilFile
func NewSmilFile(settings *helper.Settings, appName string) *SmilFile {
	s := new(SmilFile)
	s.init(settings)
	s.props["smilStreams"] = []string{}
	s.baseURI = s.host() + "/servers/" + s.serverInstance() + "/vhosts/" + s.vHostInstance() + "/applications/" + appName + "/smilfiles"
	return s
}

// Create adds the specified SMIL File configuration
func (s *SmilFile) Create(fileName string, streams []map[string]interface{}) (map[string]interface{}, error) {
	s.setRestURI(s.baseURI + "/" + fileName)
	s.props["smilStreams"] = streams

	response, err := s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, POST, "")

	return response, err
}

// Get retrieves the specified SMIL File configuration
func (s *SmilFile) Get(fileName string) (map[string]interface{}, error) {
	s.AddSkipParameter("smilStreams")
	s.setRestURI(s.baseURI + "/" + fileName)

	return s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// GetAll retrieves the list of SMIL Files for the specified Application
func (s *SmilFile) GetAll() (map[string]interface{}, error) {
	s.AddSkipParameter("smilStreams")
	s.setRestURI(s.baseURI)

	return s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// Remove deletes the specified SMIL File configuration
func (s *SmilFile) Remove(fileName string) (map[string]interface{}, error) {
	s.AddSkipParameter("smilStreams")
	s.setRestURI(s.baseURI + "/" + fileName)

	return s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, DELETE, "")
}
