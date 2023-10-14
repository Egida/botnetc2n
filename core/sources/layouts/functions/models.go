package functions

import "Nosviak2/core/sources/language/evaluator"

var (
	//stores all the functions we want to register
	//this will make sure its done correctly and safely
	Functions []evaluator.Function = make([]evaluator.Function, 0)
)

//registers the function correctly and properly
//this will make sure its done correctly and properly without issues
func RegisterFunction(r *evaluator.Function) {
	Functions = append(Functions, *r) //saves into the array correctly
}