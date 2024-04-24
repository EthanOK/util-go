package main

import (
	"fmt"
	"util-go/router"
	"util-go/tests"
)

func init() {
	fmt.Println("init()")

}
func main() {
	fmt.Println("main()")

	tests.TestParseXML()

	tests.TestParseJSON()

	r := router.Router()

	r.Run(":8090")
}
