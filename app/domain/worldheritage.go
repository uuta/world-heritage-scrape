package domain

import "encoding/xml"

type XmlData struct {
	XMLName xml.Name `xml:"query"`
	Rows    []Rows   `xml:"row"`
}

type Rows struct {
	Category string `xml:"category"`
	HttpUrl  string `xml:"http_url"`
}

type XmlBytes []byte
