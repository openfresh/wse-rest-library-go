package wserest

import (
	"strings"

	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
	"github.com/openfresh/wse-rest-library-go/entity/base"
)

// Statistics is statistics utility
type Statistics struct {
	wowza
}

// NewStatistics create Statistics object
func NewStatistics(
	settings *helper.Settings) *Statistics {
	s := new(Statistics)
	s.init(settings)
	return s
}

// GetApplicationStatistics retrieves the current Application statistics
func (s *Statistics) GetApplicationStatistics(application *Application) (map[string]interface{}, error) {
	s.setRestURI(application.baseURI + "/monitoring/current")

	return s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// GetApplicationStatisticsHistory retrieves the historic Application statistics
func (s *Statistics) GetApplicationStatisticsHistory(application *Application) (map[string]interface{}, error) {
	s.setRestURI(application.baseURI + "/monitoring/historic")

	return s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// GetIncomingApplicationStatistics retrieves the Current Incoming Stream statistics for the specifed Incoming Stream
func (s *Statistics) GetIncomingApplicationStatistics(application *Application, streamName string, appInstance string) (map[string]interface{}, error) {
	if appInstance == "" {
		appInstance = "_definst_"
	}
	s.setRestURI(application.baseURI + "/instances/" + appInstance + "/incomingstreams/" + streamName + "/monitoring/current")

	return s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// GetServerStatistics retrieves the server historical statictics
func (s *Statistics) GetServerStatistics(server *Server) (map[string]interface{}, error) {
	s.setRestURI(server.baseURI + "/monitoring/historic")

	return s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// GetServerStatisticsCurrent retrieves current statictics for the machine
func (s *Statistics) GetServerStatisticsCurrent(server *Server) (map[string]interface{}, error) {
	restURI := strings.Split("/servers/", server.baseURI)
	s.setRestURI(restURI[0] + "/machine/monitoring/current")

	return s.sendRequest(s.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}
