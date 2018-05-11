package wserest

import (
	"strconv"

	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
	"github.com/openfresh/wse-rest-library-go/entity/base"
)

// Logging is log file utility
type Logging struct {
	wowza
}

// NewLogging creates Logging object
func NewLogging(settings *helper.Settings) *Logging {
	l := new(Logging)
	l.init(settings)
	l.baseURI = l.host() + "/servers/" + l.serverInstance() + "/logfiles"
	return l
}

// GetNewestFirst retrieves the list of server log files
func (l *Logging) GetNewestFirst() (map[string]interface{}, error) {
	l.setRestURI(l.baseURI + "?order=newestFirst")

	return l.sendRequest(l.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// GetLineCount retrieves the contents of a Server Log
func (l *Logging) GetLineCount(num int) (map[string]interface{}, error) {
	l.setRestURI(l.baseURI + "/wowzastreamingengine_access.log?lineCount=" + strconv.Itoa(num))

	return l.sendRequest(l.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// Search retrieves the contents of a Server Log containing str
func (l *Logging) Search(str string) (map[string]interface{}, error) {
	l.setRestURI(l.baseURI + "/wowzastreamingengine_access.log?search=" + str)

	return l.sendRequest(l.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}
