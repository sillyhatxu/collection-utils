package collection

import (
	"reflect"
)

func (c *collection) Filter(f interface{}) *collection {
	if c == nil {
		return &collection{err: CollectionNilError}
	}
	sv, err := c.validateFilterInputSlice()
	if err != nil {
		return &collection{err: err}
	}

	funcValue, funcType, err := c.validateFilterFunc(f)
	if err != nil {
		return &collection{err: err}
	}

	resultSliceType := reflect.SliceOf(funcType.In(0))
	ret := reflect.MakeSlice(resultSliceType, 0, sv.Len())

	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i)
		if executeFilterFunc(funcValue, v) {
			ret = reflect.Append(ret, v)
		}
	}

	return &collection{
		slice: ret.Interface(),
	}
}
