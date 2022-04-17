package usecase

import (
	"encoding/xml"
	"fmt"
	"log"
	"world-heritage-scrape/app/worldheritage/domain/model"
	"world-heritage-scrape/app/worldheritage/infrastructure"
	"world-heritage-scrape/app/worldheritage/infrastructure/persistance"
	"world-heritage-scrape/config"
)

var xmlBytes model.XmlBytes
var result model.XmlData

func PutWorldHeritage() {
	fetchXml("http://whc.unesco.org/en/list/xml")
	parseXml()
	// TODO: move to main
	persistance.NewworldheritagePersistence(config.Connect())
}

func fetchXml(url string) {
	var err error
	xmlBytes, err = infrastructure.GetXml(url)
	if err != nil {
		log.Printf("Failed to get XML: %v", err)
	}
}

func parseXml() {
	xml.Unmarshal(xmlBytes, &result)
	for i := 0; i < len(result.Rows); i++ {
		fmt.Printf("Category :%v\n", result.Rows[i].Category)
	}
}
