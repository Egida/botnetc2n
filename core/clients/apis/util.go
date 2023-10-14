package apis

import (
	"encoding/json"
	"net/http"
)

//EncodeAndReturn will encode the encode input into json and write it
func EncodeAndReturn(encode interface{}, write http.ResponseWriter) {
	pure, err := json.Marshal(&encode) //marshals the information
	if err != nil {
		return
	}

	write.Write(pure) //writes the information
}

//AuthMethodToString will convert int to string
func AuthMethodToString(method int) string {
	if method == 1 {
		return "user&pass"
	} else {
		return "token"
	}
}