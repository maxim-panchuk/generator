package utils

import (
	"fmt"
	"generator/config"
	"generator/internal/generator/definitions"
	"github.com/pb33f/libopenapi/datamodel/high/base"
	"github.com/pb33f/libopenapi/orderedmap"
	"strings"
)

func UpFirst(s string) string {
	if len(s) == 0 {
		panic("upFirst panic: zero len string")
	}
	firstCharUpper := strings.ToUpper(string(s[0]))
	return firstCharUpper + s[1:]
}

func LowFirst(s string) string {
	if len(s) == 0 {
		panic("lowFirst panic: zero len string")
	}

	firstCharLower := strings.ToLower(string(s[0]))
	return firstCharLower + s[1:]
}

func GetRootFolderPath() string {
	return config.Get().PathToRepositoryRoot
}

func GetResponse(op *definitions.Operation) string {
	for _, resp := range op.Responses {
		if resp.Content != nil {
			mn := resp.Content.ModelName
			dp := LowFirst(mn)
			if resp.IsArray {
				return fmt.Sprintf("([]*%s.%sDTO, error)", dp, mn)
			} else {
				return fmt.Sprintf("(*%s.%sDTO, error)", dp, mn)
			}
		}
	}
	return "error"
}

func GetField(fieldName string, p orderedmap.Map[string, *definitions.Model]) *definitions.Model {
	model, ok := p.Get(fieldName)
	if !ok {
		panic(fmt.Sprintf("[ERROR] get field %s: no such filed in properties", fieldName))
	}
	return model
}

func GetModelDependencies(model *definitions.Model) []string {
	d := make([]string, 0)
	for pair := model.Properties.First(); pair != nil; pair = pair.Next() {
		m := pair.Value()
		if ref := m.GetReference(); ref != "" {
			defName := GetDefinitionNameFromRef(ref)
			d = append(d, LowFirst(defName))
		}
	}
	return d
}

func GetTagDependencies(tag string) []string {
	dpSet := make(map[string]struct{})
	paths := definitions.GetData().Tags[tag]
	for _, path := range paths {
		ops := path.Operations
		for _, op := range ops {
			rb := op.RequestBody
			if rb != nil && rb.Content != nil {
				dpSet[LowFirst(rb.Content.ModelName)] = struct{}{}
			}
			rspList := op.Responses
			for _, rsp := range rspList {
				if rsp.Content != nil {
					dpSet[LowFirst(rsp.Content.ModelName)] = struct{}{}
				}
			}
		}
	}
	return setToList(dpSet)
}

func GetControllerTagDependencies(tag string) []string {
	dpSet := make(map[string]struct{})
	paths := definitions.GetData().Tags[tag]
	for _, path := range paths {
		ops := path.Operations
		for _, op := range ops {
			if !op.IsTypical {
				continue
			}
			rb := op.RequestBody
			if rb != nil && rb.Content != nil {
				dpSet[LowFirst(rb.Content.ModelName)] = struct{}{}
			}
			rspList := op.Responses
			for _, rsp := range rspList {
				if rsp.Content != nil {
					dpSet[LowFirst(rsp.Content.ModelName)] = struct{}{}
				}
			}
		}
	}
	return setToList(dpSet)
}

func setToList(set map[string]struct{}) []string {
	slice := make([]string, 0, len(set))
	for key, _ := range set {
		slice = append(slice, key)
	}
	return slice
}

func GetDtoFieldType(model *definitions.Model) string {
	if model.Type == "object" {
		if ref := model.GetReference(); ref != "" {
			definitionName := GetDefinitionNameFromRef(ref)
			return "*" + LowFirst(definitionName) + "." + definitionName + "DTO"
		}
	}
	if model.Type == "array" {
		return "[]" + GetDtoFieldType(model.Items)
	}

	return ConvertToGoType(model.Type, model.Format)
}

func GetDefinitionNameFromRef(ref string) string {
	parts := strings.Split(ref, "/")
	return parts[len(parts)-1]
}

func IsArraySchema(schema *base.Schema) (isArraySchema bool, arraySchema *base.SchemaProxy) {
	if schema.Items != nil && schema.Items.N == 0 {
		return true, schema.Items.A

	}
	return false, nil
}

func ConvertToGoType(swaggerType, swaggerFormat string) string {
	switch swaggerType {
	case "integer":
		switch swaggerFormat {
		case "int32":
			return "int32"
		case "int64":
			return "int64"
		default:
			return "int64"
		}
	case "boolean":
		return "bool"
	case "string":
		switch swaggerFormat {
		case "date":
			return "time.Time"
		case "date-time":
			return "time.Time"
		default:
			return "string"
		}
	default:
		return swaggerType
	}
}

func ContainsTime(model *definitions.Model) bool {
	containsTime := false
	if model.PropKeys != nil && len(model.PropKeys) > 0 {
		for _, key := range model.PropKeys {
			pair, _ := model.Properties.Get(key)
			if ContainsTime(pair) {
				containsTime = true
			}
		}
	}
	if ConvertToGoType(model.Type, model.Format) == "time.Time" {
		containsTime = true
	}
	return containsTime
}

func ContainsPathParameters(op *definitions.Operation) bool {
	for _, p := range op.Parameters {
		if p.In == "path" {
			return true
		}
	}
	return false
}

func GetResponseByCode(code string, op *definitions.Operation) *definitions.Response {
	for _, r := range op.Responses {
		if r.Code == code {
			return r
		}
	}
	panic(fmt.Sprintf("GetResponseByCode op: %s, code: %s", op.OperationId, code))
}

func ResponseContainsSchema(op *definitions.Operation) bool {
	for _, r := range op.Responses {
		if r.Content != nil {
			return true
		}
	}
	return false
}

func ContainsEnum(model *definitions.Model) bool {
	for pair := model.Properties.First(); pair != nil; pair = pair.Next() {
		m := pair.Value()
		if m.IsEnum {
			return true
		}
	}
	return false
}

func GetModelEnums(model *definitions.Model) []*definitions.Model {
	enums := make([]*definitions.Model, 0)
	for pair := model.Properties.First(); pair != nil; pair = pair.Next() {
		m := pair.Value()
		if m.IsEnum {
			enums = append(enums, m)
		}
	}
	return enums
}

func TagContainsTypicalOperation(tag string) bool {
	paths := definitions.GetData().Tags[tag]
	for _, path := range paths {
		for _, op := range path.Operations {
			if op.IsTypical {
				return true
			}
		}
	}
	return false
}
