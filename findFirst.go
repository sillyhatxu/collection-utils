package collection

/**
f interface{} = func(v <T>) bool{}
*/
func (c *collection) FindFirst(f interface{}) *collection {
	if c == nil {
		return &collection{err: CollectionNilError}
	}
	if c.err != nil {
		return &collection{err: c.err}
	}
	sv, err := c.validateFilterInputSlice()
	if err != nil {
		return &collection{err: err}
	}
	funcValue, _, err := c.validateFilterFunc(f)
	if err != nil {
		return &collection{err: err}
	}
	if sv.Len() == 0 {
		return &collection{input: nil}
	}
	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i)
		if executeFilterFunc(funcValue, v) {
			return &collection{input: v.Interface()}
		}
	}
	return &collection{input: nil}
}
