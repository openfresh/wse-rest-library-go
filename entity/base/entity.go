package base

type Entity interface {
	RestURI() string
	EntityName() string
	SetURI(baseURI string)
	ResetURI()
}

type EntityBase struct {
	restURI string
}

func (e *EntityBase) RestURI() string {
	return e.restURI
}

func (e *EntityBase) SetRestURI(baseURI string) {
	e.restURI = baseURI
}

func (e *EntityBase) ResetURI() {
	e.restURI = ""
}
