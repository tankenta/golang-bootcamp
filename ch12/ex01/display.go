package main

import (
	"gopl.io/ch7/eval"
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	e, _ := eval.Parse("sqrt(A / pi)")
	Display("e", e)
}

func Any(value interface{}) string{
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'g', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'g', -1, 64)
	case reflect.Complex64:
		return strconv.FormatComplex(v.Complex(), 'g', -1, 64)
	case reflect.Complex128:
		return strconv.FormatComplex(v.Complex(), 'g', -1, 128)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return v.Type().String() + " value"
	}
}

func formatComposite(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		s := ""
		for i := 0; i < v.Len(); i++ {
			if i == 0 {
				s += fmt.Sprintf("[%s, ", formatComposite(v.Index(i)))
			} else if i == v.Len()-1 {
				s += fmt.Sprintf("%s]", formatComposite(v.Index(i)))
			} else {
				s += fmt.Sprintf("%s, ", formatComposite(v.Index(i)))
			}
			return s
		}
	case reflect.Struct:
		s := ""
		for i := 0; i < v.NumField(); i++ {
			if i == 0 {
				s += fmt.Sprintf("{%s: %s, ", v.Type().Field(i).Name, formatComposite(v.Field(i)))
			} else if i == v.Len()-1 {
				s += fmt.Sprintf("%s: %s}", v.Type().Field(i).Name, formatComposite(v.Field(i)))
			} else {
				s += fmt.Sprintf("%s: %s, ", v.Type().Field(i).Name, formatComposite(v.Field(i)))
			}
			return s
		}
	case reflect.Map:
		s := ""
		for i, key := range v.MapKeys() {
			if i == 0 {
				s += fmt.Sprintf("{%s: %s, ", formatComposite(key), v.MapIndex(key))
			} else if i == v.Len()-1 {
				s += fmt.Sprintf("%s: %s}", formatComposite(key), v.MapIndex(key))
			} else {
				s += fmt.Sprintf("%s: %s, ", formatComposite(key), v.MapIndex(key))
			}
			return s
		}
	}
	return formatAtom(v)
}

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatComposite(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default:
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}
