package gotype

import (
	"fmt"
	"go/token"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestOhter(t *testing.T) {
	var testpath = []string{
		"github.com/wzshiming/gotype/testdata/value",
		"github.com/wzshiming/gotype/testdata/kind",
		"github.com/wzshiming/gotype/testdata/type",
		"github.com/wzshiming/gotype/testdata/pkg",
	}
	for _, src := range testpath {
		testAll(t, src)
	}
}

func testAll(t *testing.T, src string) {
	imp := getImporter(t)
	scope, err := imp.Import(src, "")
	if err != nil {
		t.Fatal(err)
	}
	fset := imp.FileSet()
	num := scope.NumChild()
	for i := 0; i != num; i++ {
		v := scope.Child(i)
		testType(t, fset, v)
	}
}

func testType(t *testing.T, fset *token.FileSet, v Type) {

	for _, line := range strings.Split(v.Doc().Text(), "\n") {
		if line == "" {
			continue
		}
		v := v

		pos := fset.Position(v.Origin().Pos())

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
						t.Fatal(pos, "Error In num: ", to)
					}
					i, _ := strconv.ParseInt(method[1], 10, 64)
					v = v.In(int(i))
				case "Out":
					if len(method) < 2 {
						t.Fatal(pos, "Error Out num: ", to)
					}
					i, _ := strconv.ParseInt(method[1], 10, 64)
					v = v.Out(int(i))
				case "Field":
					if len(method) < 2 {
						t.Fatal(pos, "Error Field num: ", to)
					}
					i, err := strconv.ParseInt(method[1], 10, 64)
					if err != nil {
						t.Fatal(pos, "Error Field: ", err)
					}
					if v.NumField() <= int(i) {
						t.Fatal(pos, "Error Out of index range: ", to)
					}
					v = v.Field(int(i))
				case "Method":
					if len(method) < 2 {
						t.Fatal(pos, "Error Method num: ", to)
					}
					i, err := strconv.ParseInt(method[1], 10, 64)
					if err != nil {
						t.Fatal(pos, "Error Method: ", err)
					}
					if v.NumMethod() <= int(i) {
						t.Fatal(pos, "Error Out of index range: ", to)
					}
					v = v.Method(int(i))
				case "Child":
					if len(method) < 2 {
						t.Fatal(pos, "Error Child num: ", to)
					}
					i, err := strconv.ParseInt(method[1], 10, 64)
					if err != nil {
						t.Fatal(pos, "Error Child: ", err)
					}
					if v.NumMethod() <= int(i) {
						t.Fatal(pos, "Error Out of index range: ", to)
					}
					v = v.Child(int(i))
				case "FieldByName":
					if len(method) < 2 {
						t.Fatal(pos, "Error FieldByName num: ", to)
					}
					v, ok = v.FieldByName(method[1])
					if !ok {
						t.Fatal(pos, "Error not found: ", to)
					}
				case "MethodByName":
					if len(method) < 2 {
						t.Fatal(pos, "Error MethodByName num: ", to)
					}
					v, ok = v.MethodByName(method[1])
					if !ok {
						t.Fatal(pos, "Error not found: ", to)
					}
				case "ChildByName":
					if len(method) < 2 {
						t.Fatal(pos, "Error ChildByName num: ", to)
					}
					v, ok = v.ChildByName(method[1])
					if !ok {
						t.Fatal(pos, "Error not found: ", to)
					}
				default:
					t.Fatal(pos, "Error to: ", to)
				}
			}
		}
		if data, ok := tag.Lookup("Value"); ok {
			val := v.Value()
			if data != val {
				t.Fatal(pos, "Error value:", val, ":", data)
			}
		}
		if data, ok := tag.Lookup("Name"); ok {
			name := v.Name()
			if data != name {
				t.Fatal(pos, "Error name:", name, ":", data)
			}
		}
		if data, ok := tag.Lookup("String"); ok {
			str := v.String()
			if data != str {
				t.Fatal(pos, "Error string:", str, ":", data)
			}
		}
		if data, ok := tag.Lookup("Kind"); ok {
			kind := v.Kind().String()
			if data != kind {
				t.Fatal(pos, "Error kind:", kind, ":", data)
			}
		}
		if data, ok := tag.Lookup("Len"); ok {
			l := fmt.Sprint(v.Len())
			if data != l {
				t.Fatal(pos, "Error len:", l, ":", data)
			}
		}
		if data, ok := tag.Lookup("NumChild"); ok {
			num := fmt.Sprint(v.NumChild())
			if data != num {
				t.Fatal(pos, "Error num child:", num, ":", data)
			}
		}
		if data, ok := tag.Lookup("NumMethod"); ok {
			num := fmt.Sprint(v.NumMethod())
			if data != num {
				t.Fatal(pos, "Error num method:", num, ":", data)
			}
		}
		if data, ok := tag.Lookup("NumIn"); ok {
			num := fmt.Sprint(v.NumIn())
			if data != num {
				t.Fatal(pos, "Error num in:", num, ":", data)
			}
		}
		if data, ok := tag.Lookup("NumOut"); ok {
			num := fmt.Sprint(v.NumOut())
			if data != num {
				t.Fatal(pos, "Error num out:", num, ":", data)
			}
		}
		if data, ok := tag.Lookup("NumField"); ok {
			num := fmt.Sprint(v.NumField())
			if data != num {
				t.Fatal(pos, "Error num field:", num, ":", data)
			}
		}
		if data, ok := tag.Lookup("Tag"); ok {
			num := string(v.Tag())
			if data != num {
				t.Fatal(pos, "Error tag:", num, ":", data)
			}
		}
	}
}
