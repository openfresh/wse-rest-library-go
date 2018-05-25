package wserest

import (
	"strconv"

	"github.com/openfresh/wse-rest-library-go/entity/application"
	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
	"github.com/openfresh/wse-rest-library-go/entity/base"
)

// StreamFile is Stream File utility
type StreamFile struct {
	wowza
	applicationName     string
	mediaCasterType     string
	applicationInstance string
}

// NewStreamFile create StreamFile
func NewStreamFile(settings *helper.Settings, appName, streamFileName string) *StreamFile {
	s := new(StreamFile)
	s.init(settings)
	s.baseURI = s.host() + "/servers/" + s.serverInstance() + "/vhosts/" + s.vHostInstance() + "/streamfiles"

	if appName != "" {
		s.applicationName = appName
	}

	if streamFileName != "" {
		s.props["name"] = streamFileName
	}

	return s
}

// Get retrieves the specified Stream File configuration
func (s *StreamFile) Get() (map[string]interface{}, error) {
	s.AddSkipParameter("name")
	s.setRestURI(s.baseURI + "/" + s.props["name"].(string))

	return s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// GetAll retrieves the list of Stream Files for the specified VHost
func (s *StreamFile) GetAll() (map[string]interface{}, error) {
	s.AddSkipParameter("name")
	s.setRestURI(s.baseURI)

	return s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// Create adds the specified Stream File configuration
func (s *StreamFile) Create(urlProps map[string]interface{}, mediaCasterType string, applicationInstance string) (map[string]interface{}, error) {
	if mediaCasterType == "" {
		mediaCasterType = "rtp"
	}
	if applicationInstance == "" {
		applicationInstance = "_definst_"
	}
	s.mediaCasterType = mediaCasterType
	s.applicationInstance = applicationInstance
	sf := application.NewStreamFiles()
	sf.ID = "connectAppName=" + s.applicationName + "&appInstance=" + applicationInstance + "&mediaCasterType=" + mediaCasterType
	sf.Href = s.baseURI + "/streamfiles/" + sf.ID

	entities := s.getEntities([]base.Entity{sf}, "")
	s.setRestURI(s.baseURI + "/" + s.props["name"].(string))
	response, err := s.sendRequest(s.preparePropertiesForRequest(), entities, POST, "")
	if err == nil {
		items := s.getAdvancedSettings(urlProps)

		return s.addURL(items)
	}

	return response, err
}

func (s *StreamFile) addURL(advancedSettings []*helper.AdvancedSettingItem) (map[string]interface{}, error) {
	s.AddSkipParameter("name")
	s.setRestURI(s.props["restURI"].(string) + "/adv")
	s.AddAdditionalParameter("version", "1430601267443")
	s.AddAdditionalParameter("advancedSettings", advancedSettings)

	return s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, PUT, "")
}

func (s *StreamFile) getAdvancedSettings(urlProps map[string]interface{}) []*helper.AdvancedSettingItem {
	items := make([]*helper.AdvancedSettingItem, 0)
	for k, v := range urlProps {
		item := helper.NewAdvancedSettingItem()
		item.Name = k
		switch t := v.(type) {
		default:
		case bool:
			item.Value = strconv.FormatBool(t)
			item.Type = "Boolean"
		case int:
			item.Value = strconv.Itoa(t)
			item.Type = "Integer"
		case string:
			item.Value = t
			item.Type = "String"
		}
		items = append(items, item)
	}

	return items
}

// Update updates the Advanced Stream File configuration
func (s *StreamFile) Update(urlProps map[string]interface{}) (map[string]interface{}, error) {
	s.setRestURI(s.baseURI + "/" + s.props["name"].(string))
	items := s.getAdvancedSettings(urlProps)

	return s.addURL(items)
}

// Remove deletes the specified Stream File configuration
func (s *StreamFile) Remove() (map[string]interface{}, error) {
	s.AddSkipParameter("name")
	s.setRestURI(s.baseURI + "/" + s.props["name"].(string))

	return s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, DELETE, "")
}

// Connect connects
func (s *StreamFile) Connect(subFolder string) (map[string]interface{}, error) {
	s.AddSkipParameter("name")
	//	s.AddAdditionalParameter("connectAppName", s.applicationName)
	//	s.AddAdditionalParameter("appInstance", s.applicationInstance)
	//	s.AddAdditionalParameter("mediaCasterType", s.mediaCasterType)
	streamFilePath := ""
	if subFolder != "" {
		streamFilePath = subFolder + "/" + s.props["name"].(string)
	} else {
		streamFilePath = s.props["name"].(string)
	}

	baseURI := s.host() + "/servers/" + s.serverInstance() + "/vhosts/" + s.vHostInstance() + "/applications/" + s.applicationName + "/streamfiles"

	s.setRestURI(baseURI + "/" + streamFilePath + "/actions/connect")

	return s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, PUT,
		"connectAppName="+s.applicationName+"&appInstance="+s.applicationInstance+"&mediaCasterType="+s.mediaCasterType)
}

// Disconnect disconnect
func (s *StreamFile) Disconnect() (map[string]interface{}, error) {
	/*
	 * curl -X PUT --header 'Accept:application/json; charset=utf-8' --header 'Content-type:application/json; charset=utf-8'
	 * "http://localhost:8087/v2/servers/_defaultServer_/vhosts/_defaultVHost_/applications/[YOUR-APP-NAME]/instances/_definst_/incomingstreams/[STREAM-FILE-NAME]/actions/disconnectStream"
	 *
	 *
	 * "http:\/\/127.0.0.1:8087\/v2\/servers\/_defaultServer_\/vhosts\/_defaultVHost_\/applications\/live\/instances\/_definst_\/incomingstreams\/bolton_mass\/actions\/disconnectStream"
	 */
	s.AddSkipParameter("name")
	//	s.AddAdditionalParameter("connectAppName", s.applicationName)
	//	s.AddAdditionalParameter("appInstance", s.applicationInstance)
	//	s.AddAdditionalParameter("mediaCasterType", s.mediaCasterType)

	baseURI := s.host() + "/servers/" + s.serverInstance() + "/vhosts/" + s.vHostInstance() + "/applications/" + s.applicationName + "/instances/"
	s.setRestURI(baseURI + s.applicationInstance + "/incomingstreams/" + s.props["name"].(string) + ".stream/actions/disconnectStream")

	return s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, PUT, "")
}

// Reset stream
func (s *StreamFile) Reset() (map[string]interface{}, error) {
	/*
	 * curl -X PUT --header 'Accept:application/json; charset=utf-8' --header 'Content-type:application/json; charset=utf-8'
	 * "http://localhost:8087/v2/servers/_defaultServer_/vhosts/_defaultVHost_/applications/[YOUR-APP-NAME]/instances/_definst_/incomingstreams/[STREAM-FILE-NAME]/actions/resetStream"
	 *
	 *
	 * "http:\/\/127.0.0.1:8087\/v2\/servers\/_defaultServer_\/vhosts\/_defaultVHost_\/applications\/live\/instances\/_definst_\/incomingstreams\/bolton_mass\/actions\/resetStream"
	 */
	s.AddSkipParameter("name")
	baseURI := s.host() + "/servers/" + s.serverInstance() + "/vhosts/" + s.vHostInstance() + "/applications/" + s.applicationName + "/instances/"
	s.setRestURI(baseURI + s.applicationInstance + "/incomingstreams/" + s.props["name"].(string) + ".stream/actions/resetStream")

	return s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, PUT, "")
}
