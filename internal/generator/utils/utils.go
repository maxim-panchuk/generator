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
			if resp.IsArray {
				return fmt.Sprintf("([]*models.%sDTO, error)", resp.Content.ModelName)
			} else {
				return fmt.Sprintf("(*models.%sDTO, error)", resp.Content.ModelName)
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

func GetDtoFieldType(model *definitions.Model) string {
	if model.Type == "object" {
		if ref := model.GetReference(); ref != "" {
			definitionName := GetDefinitionNameFromRef(ref)
			return "*" + definitionName + "DTO"
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
