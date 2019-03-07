package bytecode

import (
	"git.exsdev.ru/ExS/monkey/ast"
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

	case *ast.InfixExpression:
		if err := c.compileInfixExpression(node); err != nil {
			return err
		}

	case *ast.IntegerLiteral:
		c.emitConst(&types.Integer{node.Value})
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
	}

	return nil
}
