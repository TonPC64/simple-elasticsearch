package main

import (
	"log"

	"github.com/TonPC64/simple-elasticsearch/elastic"
)

const EsEndPoint = ""

func main() {
	// example()
}

func example() {
	es := elastic.New(EsEndPoint)
	esReportIndex, _ := es.SetIndex("report")
	esLogIndex := es.NewIndex("log")
	esLogIndex.SearchAll("data")
	data, err := esReportIndex.SearchAll("page")
	log.Println(data, err)
}
