package bytecode

import (
	"fmt"

	"git.exsdev.ru/ExS/monkey/ast"
	"git.exsdev.ru/ExS/monkey/globals"
	"git.exsdev.ru/ExS/monkey/opcode"
	"git.exsdev.ru/ExS/monkey/types"
)

type Compiler struct {
	Bytecode
}

func New() *Compiler {
	return &Compiler{
		Bytecode{
			Code:      Instructions{},
			Constants: []types.Object{},
		},
	}
}

func (c *Compiler) Compile(node ast.Node) error {
	switch node := node.(type) {
	case *ast.Program:
		for _, s := range node.Statements {
			if err := c.Compile(s); err != nil {
				return err
			}
		}

	case *ast.ExpressionStatement:
		if err := c.Compile(node.Expression); err != nil {
			return err
		}

		c.emit(opcode.Pop)

	case *ast.InfixExpression:
		if err := c.compileInfixExpression(node); err != nil {
			return err
		}

	case *ast.PrefixExpression:
		if err := c.compilePrefixExpression(node); err != nil {
			return err
		}

	case *ast.IntegerLiteral:
		c.emitConst(&types.Integer{node.Value})

	case *ast.Boolean:
		if node.Value {
			c.emitConst(globals.True)
		} else {
			c.emitConst(globals.False)
		}

	default:
		return fmt.Errorf("unsupported instruction %T", node)
	}
	return nil
}

func (c *Compiler) emit(o opcode.Opcode, operands ...int) {
	c.Code = append(c.Code, Make(o, operands...)...)
}

func (c *Compiler) emitConst(o types.Object) {
	c.emit(opcode.Const, len(c.Constants))
	c.Constants = append(c.Constants, o)
}

func (c *Compiler) compileInfixExpression(e *ast.InfixExpression) error {
	if err := c.Compile(e.Left); err != nil {
		return err
	}
	if err := c.Compile(e.Right); err != nil {
		return err
	}

	switch e.Operator {
	case "+":
		c.emit(opcode.Add)
	case "-":
		c.emit(opcode.Sub)
	case "*":
		c.emit(opcode.Mul)
	case "/":
		c.emit(opcode.Div)
	case "<":
		c.emit(opcode.Lt)
	case ">":
		c.emit(opcode.Gt)
	case "==":
		c.emit(opcode.Eq)
	case "!=":
		c.emit(opcode.Neq)
	case "<=":
		c.emit(opcode.Leq)
	case ">=":
		c.emit(opcode.Geq)
	}

	return nil
}

func (c *Compiler) compilePrefixExpression(e *ast.PrefixExpression) error {
	if err := c.Compile(e.Right); err != nil {
		return err
	}

	switch e.Operator {
	case "!":
		c.emit(opcode.Not)
	case "-":
		c.emit(opcode.Neg)
	}

	return nil
}
