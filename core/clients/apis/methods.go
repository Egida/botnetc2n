package apis

import (
	attacks "Nosviak2/core/attack"
	"Nosviak2/core/database"
	"encoding/json"
	"net/http"

)

type Method struct {
	Header			string 			`json:"name"`
	Launched		int				`json:"attacked"`
}

//Methods will list all methods registered on cnc
func Methods(w http.ResponseWriter, r *http.Request) {
	

	//stores all the methods with the amount launched
	var amount map[int]string = make(map[int]string)

	//ranges through all methods found within the database
	for _, method := range attacks.AllMethods(make([]*attacks.Method, 0)) {
		
		//MethodSent will get all attacks with that method
		attacks, err := database.Conn.MethodSent(method.Name)
		if err != nil || len(attacks) <= 0 {
			amount[0] = method.Name; continue
		}

		//saves into the map
		amount[len(attacks)] = method.Name; continue
	}


	var times []int = make([]int, 0)
	for key := range amount {
		times = append(times, key)
	}

	//stores all the methods within the pod
	var methods []Method = make([]Method, 0)

	for _, object := range times { //saves into the array correctly allowing us to access
		methods = append(methods, Method{Header: amount[object], Launched: object})
	}

	//marshals the values without errors happening on ensures its done properly
	if target, err := json.MarshalIndent(&methods, "", "\t"); err == nil && target != nil {
		w.Write(target) //writes the objective output
	}

}