package wserest

import (
	"strconv"
	"strings"
	"time"

	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
	"github.com/openfresh/wse-rest-library-go/entity/base"
)

// DvrClipExtraction is DVR stores utility
type DvrClipExtraction struct {
	wowza
}

// NewDvrClipExtraction creates DvrClipExtraction object
func NewDvrClipExtraction(settings *helper.Settings, appName string, appInstance string) *DvrClipExtraction {
	if appInstance == "" {
		appInstance = "_definst_"
	}
	d := new(DvrClipExtraction)
	d.init(settings)
	d.baseURI = d.host() + "/servers/" + d.serverInstance() + "/vhosts/" + d.vHostInstance() + "/applications/" + appName + "/instances/" + appInstance + "/dvrstores"
	return d
}

// Create creates a new DVR store
func (d *DvrClipExtraction) Create() (map[string]interface{}, error) {
	d.setRestURI(d.baseURI)
	response, err := d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, POST, "")

	return response, err
}

// GetItem retrieves the information about a store/converter
func (d *DvrClipExtraction) GetItem(name string) (map[string]interface{}, error) {
	d.setRestURI(d.baseURI + "/" + name)

	return d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// ConvertGroup convert group
func (d *DvrClipExtraction) ConvertGroup(nameArr []string) (map[string]interface{}, error) {
	d.setNoParams()
	d.setRestURI(d.baseURI + "/actions/convert?dvrConverterStoreList=" + strings.Join(nameArr, ","))

	return d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, PUT, "")
}

// Convert converts
/*
 * /// query params
 * dvrConverterStartTime=[unix timestamp]
 * dvrConverterEndTime=[unix-timestamp]
 * dvrConverterOutputFilename=[outputfilename]
 *
 * @param $startTime is a unix timestamp
 * @param $endTime is a unix timestamp
 * @param $outputFileName is a string
 */
func (d *DvrClipExtraction) Convert(name string, startTime *time.Time, endTime *time.Time, outputFileName string) (map[string]interface{}, error) {
	d.setNoParams()
	query := ""
	if startTime != nil {
		query += "dvrConverterStartTime=" + strconv.FormatInt(startTime.Unix(), 10)
	}
	if endTime != nil {
		if query != "" {
			query += "&"
		}
		query += "dvrConverterEndTime=" + strconv.FormatInt(endTime.Unix(), 10)
	}
	if outputFileName != "" {
		if query != "" {
			query += "&"
		}
		query += "dvrConverterOutputFilename=" + outputFileName
	}
	if len(query) > 0 {
		query = "?" + query
	}

	d.setRestURI(d.baseURI + "/{$name}/actions/convert" + query)

	return d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, PUT, "")
}

// ClearCache clear cache
func (d *DvrClipExtraction) ClearCache() (map[string]interface{}, error) {
	d.setRestURI(d.baseURI + "/actions/expire")

	return d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, PUT, "")
}

// DebugConversions converts
func (d *DvrClipExtraction) DebugConversions(name string) (map[string]interface{}, error) {
	d.setRestURI(d.baseURI + "/" + name + "/actions/convert?dvrConverterDebugConversions=true")

	return d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, PUT, "")
}

// ConvertByDurationWithStartTime conver by duration with start time
func (d *DvrClipExtraction) ConvertByDurationWithStartTime(name string, startTime *time.Time, duration *time.Duration, outputFileName string) (map[string]interface{}, error) {
	d.setNoParams()
	query := ""
	if startTime != nil {
		query += "dvrConverterStartTime=" + strconv.FormatInt(startTime.Unix(), 10)
	}
	if duration != nil {
		if query != "" {
			query += "&"
		}
		query += "dvrConverterDuration=" + strconv.FormatInt(int64(*duration/time.Millisecond), 10)
	}
	if outputFileName != "" {
		if query != "" {
			query += "&"
		}
		query += "dvrConverterOutputFilename=" + outputFileName
	}
	if len(query) > 0 {
		query = "?" + query
	}

	d.setRestURI(d.baseURI + "/{$name}/actions/convert" + query)

	return d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, PUT, "")
}

// ConvertByDurationWithEndTime convert by duration with end time
func (d *DvrClipExtraction) ConvertByDurationWithEndTime(name string, endTime *time.Time, duration *time.Duration, outputFileName string) (map[string]interface{}, error) {
	d.setNoParams()
	query := ""
	if endTime != nil {
		query += "dvrConverterEndTime=" + strconv.FormatInt(endTime.Unix(), 10)
	}
	if duration != nil {
		if query != "" {
			query += "&"
		}
		query += "dvrConverterDuration=" + strconv.FormatInt(int64(*duration/time.Millisecond), 10)
	}
	if outputFileName != "" {
		if query != "" {
			query += "&"
		}
		query += "dvrConverterOutputFilename=" + outputFileName
	}
	if len(query) > 0 {
		query = "?" + query
	}

	d.setRestURI(d.baseURI + "/{$name}/actions/convert" + query)

	return d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, PUT, "")
}

// GetAll retrieves the list of DVR stores associated with this application instance
func (d *DvrClipExtraction) GetAll() (map[string]interface{}, error) {
	d.setNoParams()

	d.setRestURI(d.baseURI)

	return d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

func (d *DvrClipExtraction) setNoParams() {

}

// Remove delete DVR store
func (d *DvrClipExtraction) Remove(fileName string) (map[string]interface{}, error) {
	d.setNoParams()
	d.setRestURI(d.baseURI + "/" + fileName)

	return d.sendRequest(d.preparePropertiesForRequest(), []base.Entity{}, DELETE, "")
}
