package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		if err := json.Unmarshal(body, x); err != nil {
			return
		}
	}
}

func CheckRequestMethod(method string, w http.ResponseWriter, r *http.Request) {
	if r.Method != method {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
