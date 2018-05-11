package application

import (
	"github.com/openfresh/wse-rest-library-go/entity/base"
)

type DrmConfig struct {
	base.EntityBase
	LicenseType                           string   `json:"licenseType"`
	InUse                                 bool     `json:"inUse"`
	CupertinoEncryptionAPIBased           []string `json:"cupertinoEncryptionAPIBased"`
	EzDRMUsername                         string   `json:"ezDRMUsername"`
	EzDRMPassword                         string   `json:"ezDRMPassword"`
	BuyDRMUserKey                         string   `json:"buyDRMUserKey"`
	BuyDRMProtectSmoothStreaming          bool     `json:"buyDRMProtectSmoothStreaming"`
	BuyDRMProtectCupertinoStreaming       bool     `json:"buyDRMProtectCupertinoStreaming"`
	BuyDRMProtectMpegDashStreaming        bool     `json:"buyDRMProtectMpegDashStreaming"`
	VerimatrixProtectCupertinoStreaming   bool     `json:"verimatrixProtectCupertinoStreaming"`
	VerimatrixCupertinoKeyServerIPAddress string   `json:"verimatrixCupertinoKeyServerIpAddress"`
	VerimatrixCupertinoKeyServerPort      int      `json:"verimatrixCupertinoKeyServerPort"`
	VerimatrixCupertinoVODPerSessionKeys  bool     `json:"verimatrixCupertinoVODPerSessionKeys"`
	VerimatrixProtectSmoothStreaming      bool     `json:"verimatrixProtectSmoothStreaming"`
	VerimatrixSmoothKeyServerIPAddress    string   `json:"verimatrixSmoothKeyServerIpAddress"`
	VerimatrixSmoothKeyServerPort         int      `json:"verimatrixSmoothKeyServerPort"`
}

func NewDrmConfig() *DrmConfig {
	d := new(DrmConfig)
	d.LicenseType = "Monthly"
	d.InUse = false
	d.CupertinoEncryptionAPIBased = []string{}
	d.EzDRMUsername = ""
	d.EzDRMPassword = ""
	d.BuyDRMUserKey = ""
	d.BuyDRMProtectSmoothStreaming = false
	d.BuyDRMProtectCupertinoStreaming = false
	d.BuyDRMProtectMpegDashStreaming = false
	d.VerimatrixProtectCupertinoStreaming = false
	d.VerimatrixCupertinoKeyServerIPAddress = ""
	d.VerimatrixCupertinoKeyServerPort = 0
	d.VerimatrixCupertinoVODPerSessionKeys = false
	d.VerimatrixProtectSmoothStreaming = false
	d.VerimatrixSmoothKeyServerIPAddress = ""
	d.VerimatrixSmoothKeyServerPort = 0
	return d
}

func (d *DrmConfig) EntityName() string {
	return "drmConfig"
}

func (d *DrmConfig) SetURI(baseURI string) {
	d.SetRestURI(baseURI + "/drm")
}
