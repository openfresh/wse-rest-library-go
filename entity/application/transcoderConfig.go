package application

import (
	"github.com/openfresh/wse-rest-library-go/entity/base"
)

type TranscoderConfig struct {
	base.EntityBase
	Available         bool        `json:"available"`
	Licensed          bool        `json:"licensed"`
	Licenses          int         `json:"licenses"`
	LicensesInUse     int         `json:"licensesInUse"`
	Templates         interface{} `json:"templates"`
	TemplatesInUse    string      `json:"templatesInUse"`
	ProfileDir        string      `json:"profileDir"`
	TemplateDir       string      `json:"templateDir"`
	CreateTemplateDir []string    `json:"createTemplateDir"`
}

func NewTranscoderConfig() *TranscoderConfig {
	t := new(TranscoderConfig)
	t.Available = false
	t.Licensed = false
	t.Licenses = 0
	t.LicensesInUse = 0
	t.Templates = nil
	t.TemplatesInUse = "${SourceStreamName}.xml,transrate.xml"
	t.ProfileDir = "${com.wowza.wms.context.VHostConfigHome}/transcoder/profiles"
	t.TemplateDir = "${com.wowza.wms.context.VHostConfigHome}/transcoder/templates"
	t.CreateTemplateDir = []string{}
	return t
}

func (t *TranscoderConfig) EntityName() string {
	return "transcoderConfig"
}

func (t *TranscoderConfig) SetURI(baseURI string) {
	t.SetRestURI(baseURI + "/transcoder")
}
