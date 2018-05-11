package helper

type AdvancedSettingItem struct {
	Enabled      bool   `json:"enabled"`
	CanRemove    bool   `json:"canRemove"`
	Name         string `json:"name"`
	Value        string `json:"value"`
	DefaultValue string `json:"defaultValue"`
	Type         string `json:"type"`
	SectionName  string `json:"sectionName"`
	Section      string `json:"section"`
	Documented   bool   `json:"documented"`
}

func NewAdvancedSettingItem() *AdvancedSettingItem {
	a := new(AdvancedSettingItem)
	a.Enabled = true
	a.CanRemove = true
	a.Name = "uri"
	a.Value = ""
	a.DefaultValue = ""
	a.Type = "String"
	a.SectionName = "Common"
	a.Section = ""
	a.Documented = true
	return a
}
