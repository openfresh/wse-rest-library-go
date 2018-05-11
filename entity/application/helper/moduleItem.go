package helper

type ModuleItem struct {
	Order       int    `json:"order"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Class       string `json:"class"`
}

func NewModuleItem() *ModuleItem {
	m := new(ModuleItem)
	m.Order = 0
	m.Name = ""
	m.Description = ""
	m.Class = ""
	return m
}
