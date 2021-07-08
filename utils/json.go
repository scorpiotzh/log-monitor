package utils

import "encoding/json"

func Json(v interface{}) string {
	bys, _ := json.Marshal(v)
	return string(bys)
}
