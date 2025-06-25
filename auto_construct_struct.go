package xdi

import (
	"reflect"
)

func isAutoConstructStructDataType(dataType reflect.Type) bool {
	if dataType.Kind() == reflect.Struct {
		return isAutoConstructStruct(dataType)
	}

	if dataType.Kind() == reflect.Ptr && dataType.Elem().Kind() == reflect.Struct {
		return isAutoConstructStruct(dataType.Elem())
	}

	return false
}

func isAutoConstructStruct(dataType reflect.Type) bool {
	if dataType.Kind() != reflect.Struct {
		return false
	}

	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		_, ok := field.Tag.Lookup("inject")
		if !ok {
			return false
		}
	}

	return true
}
