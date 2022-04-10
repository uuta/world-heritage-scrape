package main

import (
	"world-heritage-scrape/app/worldheritage/infrastructure"
)

func main() {
	infrastructure.FetchXml("http://whc.unesco.org/en/list/xml")
	infrastructure.ParseXml()
}
