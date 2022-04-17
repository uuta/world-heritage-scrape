package model

import "encoding/xml"

type XmlData struct {
	XMLName xml.Name `xml:"query"`
	Rows    []Rows   `xml:"row"`
}

type Rows struct {
	Name      string  `xml:"site"`
	Category  string  `xml:"category"`
	HttpUrl   string  `xml:"http_url"`
	Latitude  float64 `xml:"latitude"`
	Longitude float64 `xml:"longitude"`
	ShortDesc string  `xml:"short_description"`
}

type Worldheritage struct {
	ID        int
	Name      string
	Category  string
	HttpUrl   string
	Latitude  float64
	Longitude float64
	ShortDesc string
}

var Worldheritages = []Worldheritage{}

type XmlBytes []byte
