package infrastructure

import (
	"encoding/xml"
	"fmt"
	"log"
	"world-heritage-scrape/app/worldheritage/entity"
)

var xmlBytes entity.XmlBytes

func FetchXml(url string) {
	var err error
	xmlBytes, err = GetXml(url)
	if err != nil {
		log.Printf("Failed to get XML: %v", err)
	}
}

func ParseXml() {
	var result entity.XmlData
	xml.Unmarshal(xmlBytes, &result)
	for i := 0; i < len(result.Rows); i++ {
		fmt.Printf("Category :%v\n", result.Rows[i].Category)
	}
}
