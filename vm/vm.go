package vm

import (
	"fmt"

	"github.com/vasilevp/monkey/bytecode"
	"github.com/vasilevp/monkey/globals"
	"github.com/vasilevp/monkey/opcode"
	"github.com/vasilevp/monkey/types"
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

		case opcode.Sub, opcode.Mul, opcode.Div, opcode.Lt, opcode.Gt, opcode.Eq, opcode.Neq, opcode.Leq, opcode.Geq:
			right := vm.pop()
			left := vm.pop()
			vm.executeBinaryOp(op, left, right)

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

func (vm *VM) executeBinaryOp(op opcode.Opcode, left, right types.Object) error {
	_, okr := right.(*types.Integer)
	_, okl := left.(*types.Integer)

	if okl && okr {
		return vm.executeBinaryIntegerOp(op, left.(*types.Integer), right.(*types.Integer))
	}

	if okl != okr {
		return fmt.Errorf("opcode %v unsupported for operand types %T %T", op, left, right)
	}

	return nil
}

func (vm *VM) executeBinaryIntegerOp(op opcode.Opcode, left, right *types.Integer) error {
	switch op {
	case opcode.Add:
		vm.push(&types.Integer{left.Value + right.Value})
	case opcode.Sub:
		vm.push(&types.Integer{left.Value - right.Value})
	case opcode.Mul:
		vm.push(&types.Integer{left.Value * right.Value})
	case opcode.Div:
		vm.push(&types.Integer{left.Value / right.Value})
	case opcode.Eq:
		vm.push(&types.Boolean{left.Value == right.Value})
	case opcode.Neq:
		vm.push(&types.Boolean{left.Value != right.Value})
	case opcode.Leq:
		vm.push(&types.Boolean{left.Value <= right.Value})
	case opcode.Geq:
		vm.push(&types.Boolean{left.Value >= right.Value})
	case opcode.Lt:
		vm.push(&types.Boolean{left.Value < right.Value})
	case opcode.Gt:
		vm.push(&types.Boolean{left.Value > right.Value})

	default:
		return fmt.Errorf("unsupported opcode %v for integer operation", op)
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
