package collection

import (
	"reflect"
)

/**
f interface{} = func(v <T>) bool{}
*/
func (c *collection) Filter(f interface{}) *collection {
	if c == nil {
		return &collection{err: CollectionNilError}
	}
	if c.err != nil {
		return &collection{err: c.err}
	}
	sv, err := c.validateSumInputSlice()
	if err != nil {
		return &collection{err: err}
	}
	funcValue, funcType, err := c.validateFilterFunc(f)
	if err != nil {
		return &collection{err: err}
	}
	resultSliceType := reflect.SliceOf(funcType.In(0))
	ret := reflect.MakeSlice(resultSliceType, 0, sv.Len())
	if sv.Len() == 0 {
		return &collection{input: ret.Interface()}
	}
	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i)
		if executeFilterFunc(funcValue, v) {
			ret = reflect.Append(ret, v)
		}
	}
	return &collection{
		input: ret.Interface(),
	}
}
