package collection

import "reflect"

/**
f interface{} = func(v <T1>) <T2>
*/
func (c *collection) Map(f interface{}) *collection {
	if c == nil {
		return &collection{err: CollectionNilError}
	}
	if c.err != nil {
		return &collection{err: c.err}
	}
	sv, err := c.validateMapInputSlice()
	if err != nil {
		return &collection{err: err}
	}

	funcValue, funcType, err := c.validateMapFunc(f)
	if err != nil {
		return &collection{err: err}
	}

	resultSliceType := reflect.SliceOf(funcType.Out(0))
	ret := reflect.MakeSlice(resultSliceType, 0, sv.Len())
	if sv.Len() == 0 {
		return &collection{input: ret.Interface()}
	}

	for i := 0; i < sv.Len(); i++ {
		v := executeMapFunc(funcValue, sv.Index(i))
		ret = reflect.Append(ret, v)
	}

	return &collection{
		input: ret.Interface(),
	}
}
