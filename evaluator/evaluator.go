package evaluator

import (
	"fmt"

	"git.exsdev.ru/ExS/monkey/ast"
	"git.exsdev.ru/ExS/monkey/types"
)

var (
	boolTrue  = &types.Boolean{Value: true}
	boolFalse = &types.Boolean{Value: false}
	NilValue  = &types.Nil{}
)

func GetBaseEnvironment() *types.Environment {
	env := types.NewEnvironment()
	env.Set("len", &types.Builtin{Fn: func(args ...types.Object) types.Object {
		if len(args) != 1 {
			return errorf("expected 1 argument, got %d", len(args))
		}

		switch arg := args[0].(type) {
		case *types.String:
			return &types.Integer{Value: int64(len(arg.Value))}
		default:
			return errorf("operator len not defined for %s", arg.Type())
		}
	}})

	return env
}

func Eval(node ast.Node, env *types.Environment) types.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node.Statements, env)
	case *ast.BlockStatement:
		return evalBlockStatement(node.Statements, env)
	case *ast.Boolean:
		return evalBooleanExpression(node.Value)
	case *ast.IntegerLiteral:
		return &types.Integer{Value: node.Value}
	case *ast.StringLiteral:
		return &types.String{Value: node.Value}
	case *ast.ArrayLiteral:
		elems := evalExpressions(node.Elements, env)
		if len(elems) == 1 && isError(elems[0]) {
			return elems[0]
		}
		return &types.Array{Elements: elems}

	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(left, node.Operator, right)
	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.ReturnStatement:
		right := Eval(node.ReturnValue, env)
		if isError(right) {
			return right
		}
		return &types.Return{Value: right}
	case *ast.LetStatement:
		if node == nil {
			return NilValue
		}
		right := Eval(node.Value, env)
		if isError(right) {
			return right
		}

		env.Set(node.Name.Value, right)

	case *ast.Identifier:
		return evalIdentifier(node, env)

	case *ast.FunctionLiteral:
		params := node.Parameters
		body := node.Body
		return &types.Function{Parameters: params, Env: env, Body: body}

	case *ast.CallExpression:
		function := Eval(node.Function, env)
		if isError(function) {
			return function
		}

		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		return applyFunction(function, args, node.Token.Literal)
	}

	return NilValue
}

func evalProgram(stmts []ast.Statement, env *types.Environment) types.Object {
	var result types.Object

	for _, s := range stmts {
		result = Eval(s, env)

		switch result := result.(type) {
		case *types.Return:
			return result.Value
		case *types.Error:
			return result
		}
	}

	return result
}

func evalBlockStatement(stmts []ast.Statement, env *types.Environment) types.Object {
	var result types.Object

	for _, s := range stmts {
		result = Eval(s, env)

		switch result := result.(type) {
		case *types.Return:
			return result
		case *types.Error:
			return result
		}
	}

	return result
}

func evalBooleanExpression(value bool) *types.Boolean {
	if value {
		return boolTrue
	}

	return boolFalse
}

func evalPrefixExpression(operator string, value types.Object) types.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(value)
	case "-":
		return evalMinusOperatorExpression(value)

	default:
		return errorf("unknown prefix operator %q", operator)
	}
}

func evalBangOperatorExpression(right types.Object) types.Object {
	if isTruthy(right) {
		return boolFalse
	}

	return boolTrue
}

func evalMinusOperatorExpression(right types.Object) types.Object {
	if right.Type() != types.IntegerT {
		return errorf("operator \"-\" not defined for %s", right.Type())
	}

	value := right.(*types.Integer)
	return &types.Integer{Value: -value.Value}
}

func evalInfixExpression(left types.Object, operator string, right types.Object) types.Object {
	switch {
	case left.Type() == types.IntegerT && right.Type() == types.IntegerT:
		return evalIntegerInfixExpression(left, operator, right)
	case left.Type() == types.StringT && right.Type() == types.StringT:
		return evalStringInfixExpression(left, operator, right)
	default:
		return errorf("operator %q not defined for (%s, %s)", operator, left.Type(), right.Type())
	}

}

func evalIntegerInfixExpression(left types.Object, operator string, right types.Object) types.Object {
	leftVal := left.(*types.Integer).Value
	rightVal := right.(*types.Integer).Value

	switch operator {
	case "+":
		return &types.Integer{leftVal + rightVal}
	case "-":
		return &types.Integer{leftVal - rightVal}
	case "*":
		return &types.Integer{leftVal * rightVal}
	case "/":
		return &types.Integer{leftVal / rightVal}

	case "<":
		return evalBooleanExpression(leftVal < rightVal)
	case ">":
		return evalBooleanExpression(leftVal > rightVal)
	case "==":
		return evalBooleanExpression(leftVal == rightVal)
	case "!=":
		return evalBooleanExpression(leftVal != rightVal)

	default:
		return errorf("unknown infix operator %q", operator)
	}
}

func evalStringInfixExpression(left types.Object, operator string, right types.Object) types.Object {
	leftVal := left.(*types.String).Value
	rightVal := right.(*types.String).Value
	switch operator {
	case "+":
		return &types.String{Value: leftVal + rightVal}
	default:
		return errorf("unknown infix operator %q", operator)
	}
}

func evalIfExpression(node *ast.IfExpression, env *types.Environment) types.Object {
	condition := Eval(node.Condition, env)

	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(node.Consequence, env)
	}

	if node.Alternative != nil {
		return Eval(node.Alternative, env)
	}

	return NilValue
}

func evalIdentifier(node *ast.Identifier, env *types.Environment) types.Object {
	val, ok := env.Get(node.Value)
	if !ok {
		return errorf("identifier not found: %q", node.Value)
	}

	return val
}

func evalExpressions(exps []ast.Expression, env *types.Environment) []types.Object {
	result := make([]types.Object, len(exps))

	for k, v := range exps {
		result[k] = Eval(v, env)
		if isError(result[k]) {
			return []types.Object{result[k]}
		}
	}

	return result
}

func applyFunction(fn types.Object, args []types.Object, literal string) types.Object {
	switch function := fn.(type) {
	case *types.Function:
		if len(args) != len(function.Parameters) {
			return errorf("%s: expected %d args, got %d", literal, len(function.Parameters), len(args))
		}
		extendedEnv := extendFunctionEnv(function, args)
		evaluated := Eval(function.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *types.Builtin:
		return function.Fn(args...)
	default:
		return errorf("not a function: %s", fn.Type())
	}
}

func extendFunctionEnv(fn *types.Function, args []types.Object) *types.Environment {
	env := types.NewEnclosedEnvironment(fn.Env)
	for k, v := range fn.Parameters {
		env.Set(v.Value, args[k])
	}

	return env
}

func unwrapReturnValue(obj types.Object) types.Object {
	if ret, ok := obj.(*types.Return); ok {
		return ret.Value
	}

	return obj
}

func isTruthy(obj types.Object) bool {
	switch obj {
	case NilValue, boolFalse:
		return false

	default:
		switch obj := obj.(type) {
		case *types.Integer:
			return obj.Value != 0
		}

		return true
	}
}

func isError(obj types.Object) bool {
	_, ok := obj.(*types.Error)
	return ok
}

func errorf(format string, args ...interface{}) *types.Error {
	return &types.Error{fmt.Errorf(format, args...)}
}
