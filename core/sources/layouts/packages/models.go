package packages

import "Nosviak2/core/sources/language/evaluator"

var (
	//stores all the collected packages
	//this will allow for better handling without issues
	Packages []evaluator.Package = make([]evaluator.Package, 0)
)

//registers the package into the array
//this will make sure its done correctly without issues
func RegisterPackage(p evaluator.Package) { //registers the package
	Packages = append(Packages, p) //saves into the array correctly
}