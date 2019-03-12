package globals

import "github.com/vasilevp/monkey/types"

// Singletons for elementary values
var (
	True  = &types.Boolean{true}
	False = &types.Boolean{false}
	Nil   = &types.Nil{}
)
