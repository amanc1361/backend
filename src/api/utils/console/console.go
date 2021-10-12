package console

import (
	"encoding/json"
	"fmt"
	"log"
)

func Preety(data interface{}) {
	b, err := json.MarshalIndent(data, "", "")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))

}
