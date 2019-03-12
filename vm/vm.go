package vm

import (
	"fmt"

	"git.exsdev.ru/ExS/monkey/bytecode"
	"git.exsdev.ru/ExS/monkey/globals"
	"git.exsdev.ru/ExS/monkey/opcode"
	"git.exsdev.ru/ExS/monkey/types"
)

type VM struct {
	stack [4096]types.Object
	sp    int
}

func New() *VM {
	return &VM{}
}

func (vm *VM) push(x types.Object) {
	vm.stack[vm.sp] = x
	vm.sp++
}

func (vm *VM) pop() types.Object {
	e := vm.stack[vm.sp-1]
	vm.sp--
	return e
}

func (vm *VM) Top() types.Object {
	return vm.stack[vm.sp-1]
}

func (vm *VM) LastPopped() types.Object {
	return vm.stack[vm.sp]
}

func (vm *VM) Run(b bytecode.Bytecode) error {
	ip := b.Code
	for len(ip) > 0 {
		op, ins, n := ip.ReadOpcode()
		ip = ip[n:]
		switch op {
		case opcode.Const:
			vm.push(b.Constants[ins[0]])

		case opcode.Pop:
			vm.pop()

		case opcode.Add:
			right := vm.pop()
			left := vm.pop()
			switch left := left.(type) {
			case *types.Integer:
				right := right.(*types.Integer)
				vm.push(&types.Integer{left.Value + right.Value})

			case *types.String:
				right := right.(*types.String)
				vm.push(&types.String{left.Value + right.Value})
			}

		case opcode.Sub:
			right := vm.pop().(*types.Integer)
			left := vm.pop().(*types.Integer)
			vm.push(&types.Integer{left.Value - right.Value})

		case opcode.Mul:
			right := vm.pop().(*types.Integer)
			left := vm.pop().(*types.Integer)
			vm.push(&types.Integer{left.Value * right.Value})

		case opcode.Div:
			right := vm.pop().(*types.Integer)
			left := vm.pop().(*types.Integer)
			vm.push(&types.Integer{left.Value / right.Value})

		case opcode.Lt:
			right := vm.pop().(*types.Integer)
			left := vm.pop().(*types.Integer)
			vm.push(&types.Boolean{left.Value < right.Value})

		case opcode.Gt:
			right := vm.pop().(*types.Integer)
			left := vm.pop().(*types.Integer)
			vm.push(&types.Boolean{left.Value > right.Value})

		case opcode.Eq:
			right := vm.pop().(*types.Integer)
			left := vm.pop().(*types.Integer)
			vm.push(&types.Boolean{left.Value == right.Value})

		case opcode.Neq:
			right := vm.pop().(*types.Integer)
			left := vm.pop().(*types.Integer)
			vm.push(&types.Boolean{left.Value != right.Value})

		case opcode.Leq:
			right := vm.pop().(*types.Integer)
			left := vm.pop().(*types.Integer)
			vm.push(&types.Boolean{left.Value <= right.Value})

		case opcode.Geq:
			right := vm.pop().(*types.Integer)
			left := vm.pop().(*types.Integer)
			vm.push(&types.Boolean{left.Value >= right.Value})

		case opcode.Neg:
			right := vm.pop().(*types.Integer)
			vm.push(&types.Integer{-right.Value})

		case opcode.Not:
			right := vm.pop()
			if !isTruthy(right) {
				vm.push(globals.True)
			} else {
				vm.push(globals.False)
			}

		default:
			return fmt.Errorf("unsupported opcode %v", op)
		}
	}

	return nil
}

func isTruthy(obj types.Object) bool {
	switch obj {
	case globals.Nil, globals.False:
		return false

	default:
		switch obj := obj.(type) {
		case *types.Integer:
			return obj.Value != 0
		}

		return true
	}
}
