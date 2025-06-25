package xdi

import (
	"errors"
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

func (c container) invokeAutoConstructStruct(dataType reflect.Type) (reflect.Value, error) {
	if !isAutoConstructStruct(dataType) {
		return reflect.Value{}, errors.New("dataType is not an auto-constructed struct")
	}

	structValue := reflect.New(dataType).Elem()
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		_, ok := field.Tag.Lookup("inject")
		if !ok {
			return reflect.Value{}, errors.New("field " + field.Name + " is not marked for injection")
		}

		fieldValue, err := c.getProviderValueByType(field.Type)
		if err != nil {
			return reflect.Value{}, err
		}

		structValue.FieldByName(field.Name).Set(fieldValue)
	}

	c.instances[dataType] = structValue
	c.providers[dataType] = structValue

	return structValue, nil
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
