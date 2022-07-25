package optional

import "reflect"

func isNil(v interface{}) bool {
	return reflect.ValueOf(v).IsNil()
}
