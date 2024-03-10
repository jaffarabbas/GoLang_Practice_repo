package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, model interface{}) {
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &model)
}
