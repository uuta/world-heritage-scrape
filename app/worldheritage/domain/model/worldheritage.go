package model

import "encoding/xml"

type XmlData struct {
	XMLName xml.Name `xml:"query"`
	Rows    []Rows   `xml:"row"`
}

type Rows struct {
	Category  string  `xml:"category"`
	HttpUrl   string  `xml:"http_url"`
	Latitude  float64 `xml:"latitude"`
	Longitude float64 `xml:"longitude"`
}

type XmlBytes []byte
