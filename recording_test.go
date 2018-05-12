package wserest

import (
	"testing"

	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
)

func TestRecording(t *testing.T) {
	// Create settings
	settings := helper.NewDefaultSettings()
	settings.SetHost(WowzaHost)
	settings.SetUseDigest(true)
	settings.SetUsername(WowzaUsername)
	settings.SetPassword(WowzaPassword)

	sf := NewRecording(settings, "", "")

	if UseWowza == "" {
		return
	}

	response, err := sf.Split("myStream")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	recordName := "myStream"
	instanceName := "_definst_"
	recorderState := "Waiting for stream"
	defaultRecorder := true
	segmentationType := "None"
	outputPath := "/usr/local/WowzaStreamingEngine/content"
	baseFile := "testme.mp4"
	fileFormat := "MP4" // or FLV
	fileVersionDelegateName := "com.wowza.wms.livestreamrecord.manager.StreamRecorderFileVersionDelegate"
	fileTemplate := "${BaseFileName}_${RecordingStartTime}_${SegmentNumber}"
	segmentDuration := 900000
	segmentSize := 10485760
	segmentSchedule := ""
	recordData := true
	startOnKeyFrame := true
	splitOnTcDiscontinuity := false
	option := "Version existing file"
	moveFirstVideoFrameToZero := true
	currentSize := 0
	currentDuration := 0
	recordingStartTime := ""

	response, err = sf.Create(recordName, instanceName, recorderState, defaultRecorder,
		segmentationType, outputPath, baseFile, fileFormat, fileVersionDelegateName, fileTemplate,
		segmentDuration, segmentSize, segmentSchedule, recordData, startOnKeyFrame, splitOnTcDiscontinuity,
		option, moveFirstVideoFrameToZero, currentSize, currentDuration, recordingStartTime)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)

	response, err = sf.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}
