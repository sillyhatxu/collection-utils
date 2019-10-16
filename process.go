package collection

import "reflect"

func executeFilterFunc(fv, v reflect.Value) bool {
	return fv.Call([]reflect.Value{v})[0].Interface().(bool)
}
