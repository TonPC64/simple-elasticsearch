# Simple Elastic

## install

```sh
go get -u github.com:TonPC64/simple-elasticsearch
```

## usage

```go
import "github.com/TonPC64/simple-elasticsearch/elastic"

func main () {
  es := elastic.New(EsEndPoint)
  esReportIndex, _ := es.SetIndex("report")
  data, err := esReportIndex.SearchAll("page")
}
```