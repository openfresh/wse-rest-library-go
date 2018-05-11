package application

import (
	"github.com/openfresh/wse-rest-library-go/entity/base"
)

type DvrConfig struct {
	base.EntityBase
	LicenseType               string `json:"licenseType"`
	InUse                     bool   `json:"inUse"`
	DvrEnable                 bool   `json:"dvrEnable"`
	WindowDuration            int    `json:"windowDuration"`
	StorageDir                string `json:"storageDir"`
	ArchiveStrategy           string `json:"archiveStrategy"`
	DvrOnlyStreaming          bool   `json:"dvrOnlyStreaming"`
	StartRecordingOnStartup   bool   `json:"startRecordingOnStartup"`
	DvrEncryptionSharedSecret string `json:"dvrEncryptionSharedSecret"`
	DvrMediaCacheEnabled      bool   `json:"dvrMediaCacheEnabled"`
	HTTPRandomizeMediaName    bool   `json:"httpRandomizeMediaName"`
}

func NewDvrConfig() *DvrConfig {
	d := new(DvrConfig)
	d.LicenseType = "Monthly"
	d.InUse = false
	d.DvrEnable = false
	d.WindowDuration = 0
	d.StorageDir = "${com.wowza.wms.context.VHostConfigHome}/dvr"
	d.ArchiveStrategy = "append"
	d.DvrOnlyStreaming = false
	d.StartRecordingOnStartup = false
	d.DvrEncryptionSharedSecret = ""
	d.DvrMediaCacheEnabled = false
	d.HTTPRandomizeMediaName = false
	return d
}

func (d *DvrConfig) EntityName() string {
	return "dvrConfig"
}

func (d *DvrConfig) SetURI(baseURI string) {
	d.SetRestURI(baseURI + "/dvr")
}
