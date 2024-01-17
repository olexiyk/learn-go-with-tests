package main

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	walkFn := func(v reflect.Value) {
		walk(v.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkFn(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkFn(val.Index(i))
		}
	case reflect.Chan:
		for {
			v, ok := val.Recv()
			if ok {
				walkFn(v)
			} else {
				break
			}
		}
	case reflect.Map:
		iter := val.MapRange()
		for iter.Next() {
			walkFn(iter.Value())
		}
	case reflect.Func:
		values := val.Call(nil)
		for _, v := range values {
			walkFn(v)
		}
	}
}
