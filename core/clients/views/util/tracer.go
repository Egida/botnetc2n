package util

import (
	"Nosviak2/core/database"
)

//runs the parent tracer properly
//this will ensure its done without any errors
func ParentTracer(socket int, src []int) ([]int, error) {

	//sets the socket guide
	//allows for proper control
	var p int = socket
	

	//starts looping through properly
	//allows for secure usage without issues
	for { //starts the looping properly
		user, err := database.Conn.GetUserViaParent(p)
		if err != nil || user == nil {
			return src, nil
		}
		
		if user.Parent == p {
			//checks if the issue is that they match
			if user.Parent == p {src=append(src, user.Identity)}
			break //stops the looping properly
		}

		src = append(src, user.Identity)

		//sets the parent
		p = user.Parent
	}
	

	return src, nil
}

//searchs the tracer properly
//this will try to make sure its done without any errors
func SearchTracer(tracer []int, want int) bool { //returns boolean
	//ranges through the tracer
	//this will ensure its found if its there
	for _, object := range tracer { //ranges through
		if object == want {return true} //sets true
	}; return false //returns false properly
}