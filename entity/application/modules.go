package application

import (
	"github.com/openfresh/wse-rest-library-go/entity/application/helper"
	"github.com/openfresh/wse-rest-library-go/entity/base"
)

type Modules struct {
	base.EntityBase
	ModuleList []*helper.ModuleItem `json:"moduleList"`
}

func NewModules() *Modules {
	m := new(Modules)
	m.ModuleList = []*helper.ModuleItem{}
	m.ModuleList = append(m.ModuleList, m.GetModuleItem("base", "Base", "com.wowza.wms.module.ModuleCore", -1))
	m.ModuleList = append(m.ModuleList, m.GetModuleItem("logging", "Client Logging", "com.wowza.wms.module.ModuleClientLogging", -1))
	m.ModuleList = append(m.ModuleList, m.GetModuleItem("flvplayback", "FLVPlayback", "com.wowza.wms.module.ModuleFLVPlayback", -1))
	return m
}

func (m *Modules) GetModuleItem(name string, description string, class string, order int) *helper.ModuleItem {
	if order < 0 {
		order = len(m.ModuleList)
	}

	mi := helper.NewModuleItem()
	mi.Order = order
	mi.Name = name
	mi.Description = description
	mi.Class = class

	return mi
}

func (m *Modules) EntityName() string {
	return "modules"
}

func (m *Modules) SetURI(baseURI string) {
	m.SetRestURI(baseURI + "/streamconfiguration")
}
