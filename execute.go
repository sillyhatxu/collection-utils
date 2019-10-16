package collection

import (
	"reflect"
	"sort"
)

func executeFilterFunc(fv, v reflect.Value) bool {
	return fv.Call([]reflect.Value{v})[0].Interface().(bool)
}

func executeMapFunc(fv, arg reflect.Value) reflect.Value {
	return fv.Call([]reflect.Value{arg})[0]
}

func executeSort(fv, ret reflect.Value) {
	less := func(i, j int) bool {
		return fv.Call([]reflect.Value{ret.Index(i), ret.Index(j)})[0].Interface().(bool)
	}
	sort.Sort(&customFunc{length: ret.Len(), less: less, swap: reflect.Swapper(ret.Interface())})
}
