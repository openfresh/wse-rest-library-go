package wserest

import (
	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
	"github.com/openfresh/wse-rest-library-go/entity/base"
)

// Recording is Stream Recorder utility
type Recording struct {
	wowza
}

// NewRecording creates Recording object
func NewRecording(settings *helper.Settings, appName string, appInstance string) *Recording {
	if appName == "" {
		appName = "live"
	}
	if appInstance == "" {
		appInstance = "_definst_"
	}
	r := new(Recording)
	r.init(settings)
	r.props["recorderName"] = "myStream"
	r.props["instanceName"] = "_definst_"
	r.props["recorderState"] = "Waiting for stream"
	r.props["defaultRecorder"] = true
	r.props["segmentationType"] = "None"
	r.props["outputPath"] = ""
	r.props["baseFile"] = "myrecord.mp4"
	r.props["fileFormat"] = "MP4"
	r.props["fileVersionDelegateName"] = "com.wowza.wms.livestreamrecord.manager.StreamRecorderFileVersionDelegate"
	r.props["fileTemplate"] = "${BaseFileName}_${RecordingStartTime}_${SegmentNumber}"
	r.props["segmentDuration"] = 900000
	r.props["segmentSize"] = 10485760
	r.props["segmentSchedule"] = "0 * * * * *"
	r.props["recordData"] = true
	r.props["startOnKeyFrame"] = true
	r.props["splitOnTcDiscontinuity"] = false
	r.props["option"] = "Version existing file"
	r.props["moveFirstVideoFrameToZero"] = true
	r.props["currentSize"] = 0
	r.props["currentDuration"] = 0
	r.props["recordingStartTime"] = ""
	r.baseURI = r.host() + "/servers/" + r.serverInstance() + "/vhosts/" + r.vHostInstance() + "/applications/" + appName + "/instances/" + appInstance + "/streamrecorders"
	return r
}

// Create creates a new Stream Recorder in the specified Application Instance and starts recording
func (r *Recording) Create(
	recorderName string,
	instanceName string,
	recorderState string,
	defaultRecorder bool,
	segmentationType string,
	outputPath string,
	baseFile string,
	fileFormat string,
	fileVersionDelegateName string,
	fileTemplate string,
	segmentDuration int,
	segmentSize int,
	segmentSchedule string,
	recordData bool,
	startOnKeyFrame bool,
	splitOnTcDiscontinuity bool,
	option string,
	moveFirstVideoFrameToZero bool,
	currentSize int,
	currentDuration int,
	recordingStartTime string,
) (map[string]interface{}, error) {
	r.props["recorderName"] = recorderName
	r.props["instanceName"] = instanceName
	r.props["recorderState"] = recorderState
	r.props["defaultRecorder"] = defaultRecorder
	r.props["segmentationType"] = segmentationType
	r.props["outputPath"] = outputPath
	r.props["baseFile"] = baseFile
	r.props["fileFormat"] = fileFormat
	r.props["fileVersionDelegateName"] = fileVersionDelegateName
	r.props["fileTemplate"] = fileTemplate
	r.props["segmentDuration"] = segmentDuration
	r.props["segmentSize"] = segmentSize
	r.props["segmentSchedule"] = segmentSchedule
	r.props["recordData"] = recordData
	r.props["startOnKeyFrame"] = startOnKeyFrame
	r.props["splitOnTcDiscontinuity"] = splitOnTcDiscontinuity
	r.props["option"] = option
	r.props["moveFirstVideoFrameToZero"] = moveFirstVideoFrameToZero
	r.props["currentSize"] = currentSize
	r.props["currentDuration"] = currentDuration
	r.props["recordingStartTime"] = recordingStartTime

	r.setRestURI(r.baseURI)
	response, err := r.sendRequest(r.preparePropertiesForRequest(), []base.Entity{}, POST, "")

	return response, err
}

// GetAll retrieves the list of Stream Recorders
func (r *Recording) GetAll() (map[string]interface{}, error) {
	r.setNoParams()

	r.setRestURI(r.baseURI)
	return r.sendRequest(r.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// GetRecorder retrieves the specifed Stream Recorder
func (r *Recording) GetRecorder(recorderName string) (map[string]interface{}, error) {
	r.setRestURI(r.baseURI + "/" + recorderName)
	r.setNoParams()

	return r.sendRequest(r.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// GetDefaultParams retrieves a Stream Recorder of the requested name, popluated with the default values
func (r *Recording) GetDefaultParams(recorderName string) (map[string]interface{}, error) {
	r.setRestURI(r.baseURI + "/" + recorderName + "/default")
	r.setNoParams()

	return r.sendRequest(r.preparePropertiesForRequest(), []base.Entity{}, GET, "")
}

// Stop stop recording
func (r *Recording) Stop(recorderName string) (map[string]interface{}, error) {
	r.setRestURI(r.baseURI + "/" + recorderName + "/actions/stopRecording")
	r.setNoParams()

	return r.sendRequest(r.preparePropertiesForRequest(), []base.Entity{}, PUT, "")
}

// Split splits recording
func (r *Recording) Split(recorderName string) (map[string]interface{}, error) {
	r.setRestURI(r.baseURI + "/" + recorderName + "/actions/splitRecording")
	r.setNoParams()

	return r.sendRequest(r.preparePropertiesForRequest(), []base.Entity{}, PUT, "")
}

func (r *Recording) setNoParams() {
	r.AddSkipParameter("recordName")
	r.AddSkipParameter("instanceName")
	r.AddSkipParameter("recorderState")
	r.AddSkipParameter("defaultRecorder")
	r.AddSkipParameter("segmentationType")
	r.AddSkipParameter("outputPath")
	r.AddSkipParameter("baseFile")
	r.AddSkipParameter("fileFormat")
	r.AddSkipParameter("fileVersionDelegateName")
	r.AddSkipParameter("fileTemplate")
	r.AddSkipParameter("segmentDuration")
	r.AddSkipParameter("segmentSize")
	r.AddSkipParameter("segmentSchedule")
	r.AddSkipParameter("startOnKeyFrame")
	r.AddSkipParameter("splitOnTcDiscontinuity")
	r.AddSkipParameter("option")
	r.AddSkipParameter("moveFirstVideoFrameToZero")
	r.AddSkipParameter("currentSize")
	r.AddSkipParameter("currentDuration")
	r.AddSkipParameter("recordingStartTime")
}
