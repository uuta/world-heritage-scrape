package usecase

import (
	"encoding/xml"
	"fmt"
	"log"
	"world-heritage-scrape/app/worldheritage/domain/model"
	"world-heritage-scrape/app/worldheritage/infrastructure"
)

var xmlBytes model.XmlBytes

func PutWorldHeritage() {
	fetchXml("http://whc.unesco.org/en/list/xml")
	parseXml()
}

func fetchXml(url string) {
	var err error
	xmlBytes, err = infrastructure.GetXml(url)
	if err != nil {
		log.Printf("Failed to get XML: %v", err)
	}
}

func parseXml() {
	var result model.XmlData
	xml.Unmarshal(xmlBytes, &result)
	for i := 0; i < len(result.Rows); i++ {
		fmt.Printf("Category :%v\n", result.Rows[i].Category)
	}
}
