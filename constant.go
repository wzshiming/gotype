package gotype

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/token"
	"strconv"
)

func constantEval(expr ast.Node, iota int64, info *infoFile) (r constant.Value, err error) {
	switch t := expr.(type) {
	case *ast.BasicLit:
		switch t.Kind {
		case token.STRING:
			r = constant.MakeString(t.Value[1 : len(t.Value)-1])
			return r, nil
		case token.INT:
			i, err := strconv.ParseInt(t.Value, 0, 0)
			if err != nil {
				return nil, err
			}
			r = constant.MakeInt64(i)
			return r, nil
		case token.FLOAT:
			i, err := strconv.ParseFloat(t.Value, 0)
			if err != nil {
				return nil, err
			}
			r = constant.MakeFloat64(i)
			return r, nil
		}
	case *ast.UnaryExpr:
		x, err := constantEval(t.X, iota, info)
		if err != nil {
			return nil, err
		}
		r, err = constantUnaryOp(t.Op, x)
		if err != nil {
			return nil, err
		}
		return r, nil
	case *ast.BinaryExpr:
		x, err := constantEval(t.X, iota, info)
		if err != nil {
			return nil, err
		}
		y, err := constantEval(t.Y, iota, info)
		if err != nil {
			return nil, err
		}
		r, err = constantBinaryOp(t.Op, x, y)
		if err != nil {
			return nil, err
		}
		return r, nil
	case *ast.Ident:
		if t.Name == "iota" {
			r = constant.MakeInt64(iota)
			return r, nil
		} else if val, ok := info.GetPkgOrType(t.Name); ok && val.Kind() == Declaration {
			val = val.Declaration()
			switch val.Kind() {
			case Int, Int8, Int16, Int32, Int64:
				i, err := strconv.ParseInt(val.Value(), 0, 0)
				if err != nil {
					return nil, err
				}
				r = constant.MakeInt64(i)
				return r, nil
			case Uint, Uint8, Uint16, Uint32, Uint64:
				i, err := strconv.ParseUint(val.Value(), 0, 0)
				if err != nil {
					return nil, err
				}
				r = constant.MakeUint64(i)
				return r, nil
			case Float32, Float64:
				i, err := strconv.ParseFloat(val.Value(), 0)
				if err != nil {
					return nil, err
				}
				r = constant.MakeFloat64(i)
				return r, nil
			case String:
				str := val.Value()
				r = constant.MakeString(str[1 : len(str)-1])
				return r, nil
			default:
				// iota
				i, err := strconv.ParseInt(val.Value(), 0, 0)
				if err == nil {
					r = constant.MakeInt64(i)
					return r, nil
				}
			}
		} else {
			return nil, fmt.Errorf("undefined ident")
		}
	case *ast.ParenExpr:
		r, err = constantEval(t.X, iota, info)
		if err != nil {
			return nil, err
		}
		return r, nil
	case *ast.CallExpr:
		if len(t.Args) == 0 {
			return nil, fmt.Errorf("undefined call")
		}
		r, err = constantEval(t.Args[0], iota, info)
		if err != nil {
			return nil, err
		}
		return r, nil
	default:
		return nil, fmt.Errorf("undefined expr")
	}
	return constant.MakeUnknown(), nil
}

func constantUnaryOp(op token.Token, y constant.Value) (r constant.Value, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("%v", x)
		}
	}()

	r = constant.UnaryOp(op, y, 0)
	return r, nil
}

func constantBinaryOp(op token.Token, x, y constant.Value) (r constant.Value, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("%v", x)
		}
	}()
	switch op {
	case token.SHL, token.SHR:
		n, _ := constant.Uint64Val(y)
		r = constant.Shift(x, op, uint(n))
	case token.EQL, token.NEQ, token.LSS, token.LEQ, token.GTR, token.GEQ:
		r = constant.MakeBool(constant.Compare(x, op, y))
	default:
		r = constant.BinaryOp(x, op, y)
	}
	return r, nil
}
