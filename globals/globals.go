package globals

import "git.exsdev.ru/ExS/monkey/types"

var (
	True  = &types.Boolean{Value: true}
	False = &types.Boolean{Value: false}
	Nil   = &types.Nil{}
)
