package optional

type Optional interface {
	Map(mapping func(old interface{}) (v interface{})) Optional
	OrElse(v interface{}) interface{}
}

type optional struct {
	v interface{}
}

func (o *optional) OrElse(v interface{}) interface{} {

	if isNil(o.v) {
		return v
	}

	return o.v
}

func (o *optional) Map(mapping func(old interface{}) (v interface{})) Optional {
	if isNil(o.v) {
		return o
	}

	v := mapping(o.v)
	if op, ok := v.(Optional); ok {
		return op
	}

	return &optional{
		v: v,
	}
}

func OfNullable(v interface{}) Optional {

	return &optional{
		v: v,
	}
}

func Of(v interface{}, err error) Optional {

	// 有异常  崩溃
	if err != nil {
		panic(err)
	}

	return &optional{
		v: v,
	}
}
