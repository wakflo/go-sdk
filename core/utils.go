package core

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(data interface{}) {
	prettyJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return
	}
	fmt.Printf("%s\n", prettyJSON)
}
