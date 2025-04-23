package reflection

import "reflect"

func Runs(x interface{}, fn func(input string)) {
	value := getValue(x)

	runsValue := func(value reflect.Value) {
		Runs(value.Interface(), fn)
	}

	switch value.Kind() {
	case reflect.String:
		fn(value.String())
	case reflect.Struct:
		for i := range value.NumField() {
			runsValue(value.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := range value.Len() {
			runsValue(value.Index(i))
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			runsValue(value.MapIndex(key))
		}
	}
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	return value
}
