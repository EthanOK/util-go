package tests

import (
	"fmt"
	"os"
	"util-go/utils"
)

func TestParseJSON() {
	data_bytes, err := os.ReadFile("files/data.json")
	if err != nil {
		panic(err)
	}
	var data map[string]interface{}
	err = utils.ParseJson(data_bytes, &data)
	if err != nil {
		// handle error
		panic(err)
	}
	// do something with data
	fmt.Println(data)

}
