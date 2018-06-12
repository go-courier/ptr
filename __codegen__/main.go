package main

import (
	"github.com/go-courier/codegen"
)

func main() {
	basicTypes := []codegen.BuiltInType{
		"bool",

		"int",
		"int8",
		"int16",
		"int32",
		"int64",

		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"uintptr",

		"float32",
		"float64",
		"complex64",
		"complex128",

		"string",
		"byte",
		"rune",
	}

	{
		file := codegen.NewFile("ptr", "ptr.go")

		for _, basicType := range basicTypes {
			file.WriteBlock(
				codegen.Func(codegen.Var(basicType, "v")).
					Return(codegen.Var(codegen.Star(basicType))).
					Named(codegen.UpperCamelCase(string(basicType))).
					Do(codegen.Return(codegen.Unary(codegen.Id("v")))),
			)
			file.WriteRune('\n')
		}

		file.WriteFile()
	}

	{
		file := codegen.NewFile("ptr", "ptr_test.go")

		for _, basicType := range basicTypes {
			name := codegen.UpperCamelCase(string(basicType))

			funcType := codegen.Func(codegen.Var(codegen.Star(codegen.Type(file.Use("testing", "T"))), "t")).
				Named("Test" + name)

			switch basicType {
			case "string":
				file.WriteBlock(funcType.Do(
					codegen.Call(
						file.Use("github.com/stretchr/testify/assert", "Equal"),
						codegen.Id("t"),
						codegen.Val("string"),
						codegen.Expr("*?", codegen.Call(name, codegen.Val("string"))),
					),
				))
			case "rune":
				file.WriteBlock(funcType.Do(
					codegen.Call(
						file.Use("github.com/stretchr/testify/assert", "Equal"),
						codegen.Id("t"),
						codegen.Val('r'),
						codegen.Expr("*?", codegen.Call(name, codegen.Val('r'))),
					),
				))
			case "byte":
				v := codegen.Expr("?[0]", codegen.Val([]byte("bytes")))

				file.WriteBlock(funcType.Do(
					codegen.Call(
						file.Use("github.com/stretchr/testify/assert", "Equal"),
						codegen.Id("t"),
						v,
						codegen.Expr("*?", codegen.Call(name, v)),
					),
				))
			case "bool":
				file.WriteBlock(funcType.Do(
					codegen.Call(
						file.Use("github.com/stretchr/testify/assert", "Equal"),
						codegen.Id("t"),
						codegen.Val(true),
						codegen.Expr("*?", codegen.Call(name, codegen.Val(true))),
					),
					codegen.Call(
						file.Use("github.com/stretchr/testify/assert", "Equal"),
						codegen.Id("t"),
						codegen.Val(false),
						codegen.Expr("*?", codegen.Call(name, codegen.Val(false))),
					),
				))
			default:
				file.WriteBlock(funcType.Do(
					codegen.Call(
						file.Use("github.com/stretchr/testify/assert", "Equal"),
						codegen.Id("t"),
						codegen.CallWith(basicType, codegen.Val(1)),
						codegen.Expr("*?", codegen.Call(name, codegen.Val(1))),
					),
				))
			}

			file.WriteRune('\n')
		}

		file.WriteFile()
	}
}
