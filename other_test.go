package gotype

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestOhter(t *testing.T) {
	var testpath = []string{
		"github.com/wzshiming/gotype/testdata/value",
		"github.com/wzshiming/gotype/testdata/kind",
		"github.com/wzshiming/gotype/testdata/struct",
	}
	for _, src := range testpath {
		testAll(t, Import(t, src))
	}
}

func testAll(t *testing.T, scope Type) {
	num := scope.NumChild()
	for i := 0; i != num; i++ {
		v := scope.Child(i)
		testType(t, v)
	}
}

func testType(t *testing.T, v Type) {

	for _, line := range strings.Split(v.Doc().Text(), "\n") {
		if line == "" {
			continue
		}
		v := v
		tag := reflect.StructTag(line)
		if data, ok := tag.Lookup("To"); ok {
			st := strings.Split(data, ",")
			for _, to := range st {
				method := strings.Split(to, ":")
				switch method[0] {
				case "Elem":
					v = v.Elem()
				case "Declaration":
					v = v.Declaration()
				case "Key":
					v = v.Key()
				case "In":
					if len(method) < 2 {
						t.Fatal("Error In num: ", to)
					}
					i, _ := strconv.ParseInt(method[1], 10, 64)
					v = v.In(int(i))
				case "Out":
					if len(method) < 2 {
						t.Fatal("Error Out num: ", to)
					}
					i, _ := strconv.ParseInt(method[1], 10, 64)
					v = v.Out(int(i))
				case "Field":
					if len(method) < 2 {
						t.Fatal("Error Field num: ", to)
					}
					i, err := strconv.ParseInt(method[1], 10, 64)
					if err != nil {
						t.Fatal("Error Field: ", err)
					}
					if v.NumField() <= int(i) {
						t.Fatal("Error Out of index range: ", to)
					}
					v = v.Field(int(i))
				case "Method":
					if len(method) < 2 {
						t.Fatal("Error Method num: ", to)
					}
					i, err := strconv.ParseInt(method[1], 10, 64)
					if err != nil {
						t.Fatal("Error Method: ", err)
					}
					if v.NumMethod() <= int(i) {
						t.Fatal("Error Out of index range: ", to)
					}
					v = v.Method(int(i))
				case "Child":
					if len(method) < 2 {
						t.Fatal("Error Child num: ", to)
					}
					i, err := strconv.ParseInt(method[1], 10, 64)
					if err != nil {
						t.Fatal("Error Child: ", err)
					}
					if v.NumMethod() <= int(i) {
						t.Fatal("Error Out of index range: ", to)
					}
					v = v.Child(int(i))
				case "FieldByName":
					if len(method) < 2 {
						t.Fatal("Error FieldByName num: ", to)
					}
					v, ok = v.FieldByName(method[1])
					if !ok {
						t.Fatal("Error not found: ", to)
					}
				case "MethodByName":
					if len(method) < 2 {
						t.Fatal("Error MethodByName num: ", to)
					}
					v, ok = v.MethodByName(method[1])
					if !ok {
						t.Fatal("Error not found: ", to)
					}
				case "ChildByName":
					if len(method) < 2 {
						t.Fatal("Error ChildByName num: ", to)
					}
					v, ok = v.ChildByName(method[1])
					if !ok {
						t.Fatal("Error not found: ", to)
					}
				default:
					t.Fatal("Error to: ", to)
				}
			}
		}
		if data, ok := tag.Lookup("Value"); ok {
			val := v.Value()
			if data != val {
				t.Fatal("Error value:", val, data)
			}
		}
		if data, ok := tag.Lookup("Name"); ok {
			name := v.Name()
			if data != name {
				t.Fatal("Error name:", name, data)
			}
		}
		if data, ok := tag.Lookup("String"); ok {
			str := v.String()
			if data != str {
				t.Fatal("Error string:", str, data)
			}
		}
		if data, ok := tag.Lookup("Kind"); ok {
			kind := v.Kind().String()
			if data != kind {
				t.Fatal("Error kind:", kind, data)
			}
		}
		if data, ok := tag.Lookup("Len"); ok {
			l := fmt.Sprint(v.Len())
			if data != l {
				t.Fatal("Error len:", l, data)
			}
		}
		if data, ok := tag.Lookup("NumMethod"); ok {
			num := fmt.Sprint(v.NumMethod())
			if data != num {
				t.Fatal("Error num method:", num, data)
			}
		}
		if data, ok := tag.Lookup("NumIn"); ok {
			num := fmt.Sprint(v.NumIn())
			if data != num {
				t.Fatal("Error num in:", num, data)
			}
		}
		if data, ok := tag.Lookup("NumOut"); ok {
			num := fmt.Sprint(v.NumOut())
			if data != num {
				t.Fatal("Error num out:", num, data)
			}
		}
		if data, ok := tag.Lookup("NumField"); ok {
			num := fmt.Sprint(v.NumField())
			if data != num {
				t.Fatal("Error num field:", num, data)
			}
		}
		if data, ok := tag.Lookup("Tag"); ok {
			num := string(v.Tag())
			if data != num {
				t.Fatal("Error tag:", num, data)
			}
		}
	}
}
