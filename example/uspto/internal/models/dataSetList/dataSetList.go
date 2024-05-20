package dataSetList

import ()

type dataSetListDTO struct {
	Total int64    `json:"total"`
	Apis  []object `json:"apis"`
}
