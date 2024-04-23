package models

import "encoding/xml"

type Peoples struct {
	XMLName xml.Name `xml:"peoples"`
	Version string   `xml:"version,attr"`
	Peops   []People `xml:"people"`
}

type People struct {
	XMLName xml.Name `xml:"people"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Gender  string   `xml:"gender"`
}
