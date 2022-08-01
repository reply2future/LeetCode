package utils

import (
	"encoding/json"
    "io/ioutil"
	"fmt"
)

func ReadJsonFromFile[T any](path string) T {
	content, err := ioutil.ReadFile(path)
    if err != nil {
        panic(fmt.Sprintf("Error when opening file: %v", err))
    }
 
    var payload T
    err = json.Unmarshal(content, &payload)
    if err != nil {
        panic(fmt.Sprintf("Error during Unmarshal(): %v", err))
    }
	return payload 
}