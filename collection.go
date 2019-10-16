package collection

type collection struct {
	input interface{}
	err   error
}

func Stream(input interface{}) *collection {
	return &collection{input: input}
}

func (c *collection) Result() (interface{}, error) {
	return c.input, c.err
}
