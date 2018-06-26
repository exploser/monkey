package types

import (
	"fmt"
	"strings"

	"git.exsdev.ru/ExS/monkey/ast"
)

var _ Object = new(Function)

const FunctionT ObjectType = "Function"

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (*Function) Type() ObjectType {
	return FunctionT
}

func (i *Function) String() string {
	params := make([]string, len(i.Parameters))
	for k, v := range i.Parameters {
		params[k] = v.String()
	}

	return fmt.Sprintf("func(%s) %s", strings.Join(params, ", "), i.Body)
}
