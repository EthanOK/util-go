package main

import (
	"fmt"
	"util-go/tests"
)

func init() {
	fmt.Println("init()")

}
func main() {
	fmt.Println("main()")

	tests.TestParseXML()

	tests.TestParseJSON()

}
