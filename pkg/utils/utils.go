package utils

import(

	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, model interface{}){

	if body, err := ioutil.ReadAll(r.Body); err == nil{

		if err := json.Unmarshal([]byte(body), model); err != nil{

			return
		}
	}

}