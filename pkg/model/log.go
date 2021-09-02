package model

type Object = map[string]interface{}

type Entry struct {
	Origin          string
	ParsedObject    Object
	PipelinedObject Object
	Err             error
}

func NewEntry(origin string) *Entry {
	return &Entry{
		Origin:          origin,
		ParsedObject:    make(Object),
		PipelinedObject: make(Object),
		Err:             nil,
	}
}
