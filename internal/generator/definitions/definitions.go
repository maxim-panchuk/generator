package definitions

import (
	"github.com/pb33f/libopenapi/orderedmap"
	"sync"
)

type Model struct {
	ModelName   string
	Description string
	Type        string
	Format      string
	PropKeys    []string
	Properties  *orderedmap.Map[string, *Model]
	Items       *Model
	Ref         string
}

func (m *Model) GetReference() (ref string) {
	if m.Ref != "" {
		return m.Ref
	}
	if m.Items != nil {
		return m.Items.GetReference()
	}
	return
}

type Models map[string]*Model

type Parameter struct {
	Name        string
	Description string
	In          string
	Schema      *Model
}

type Response struct {
	Code        string
	Description string
	Content     *Model
}

func (r *Response) IsDefault() bool {
	return r.Code == "default"
}

type XMeta struct {
	Object string
	Model  *Model
	Type   string
}

func (xm *XMeta) IncludeModel(name string) {
	xm.Model = GetData().Models[name]
}

type Operation struct {
	Tag         string
	Summary     string
	Description string
	OperationId string
	Parameters  []*Parameter
	RequestBody *Model
	Responses   []*Response
	IsTypical   bool
	XMeta       *XMeta
}

type Path struct {
	Url        string
	Operations []*Operation
}

type Tags map[string][]*Path

type Data struct {
	Models Models
	Tags   Tags
}

var once sync.Once
var data *Data

func GetData() *Data {
	once.Do(func() {
		data = &Data{
			Models: make(Models),
			Tags:   make(Tags),
		}
	})
	return data
}
