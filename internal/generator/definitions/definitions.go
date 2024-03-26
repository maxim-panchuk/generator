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
}

type Models map[string]*Model

type Data struct {
	Models Models
}

var once sync.Once
var data *Data

func GetData() *Data {
	once.Do(func() {
		data = &Data{
			Models: make(Models),
		}
	})
	return data
}
