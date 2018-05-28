package elastic

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type ElasticClient struct {
	EndPoint string
}

func New(url string) ElasticClient {
	return ElasticClient{
		EndPoint: url,
	}
}

type elasticIndex struct {
	EndPoint string
	Index    string
}

func (e *ElasticClient) SetIndex(Index string) (elasticIndex, error) {
	if e.EndPoint == "" {
		fmt.Println("Don't have endpoint.")
		return elasticIndex{}, errors.New("Don't have endpoint.\n")
	} else {
		return elasticIndex{
			EndPoint: e.EndPoint,
			Index:    Index,
		}, nil
	}
}

// elasticIndex Func

func (e *elasticIndex) Insert(Type string, Data interface{}) (ElasticResPost, error) {
	return e.InsertWithID(Type, "", Data)
}

func (e *elasticIndex) InsertWithID(Type, ID string, Data interface{}) (ElasticResPost, error) {

	if e.EndPoint != "" {

		byteData, _ := json.Marshal(Data)
		postBody := bytes.NewReader(byteData)
		esHttp := elasticHttp{
			EndPoint: e.EndPoint,
			Index:    e.Index,
			Type:     Type,
			ID:       ID,
			Path:     "/" + e.Index + "/" + Type + "/" + ID,
		}

		res, err := esHttp.POST(postBody)
		resBody, _ := ioutil.ReadAll(res.Body)
		var resData ElasticResPost
		json.Unmarshal(resBody, &resData)

		if !resData.Created {
			return ElasticResPost{}, errors.New("Insert failed.")
		}
		return resData, err
	}
	fmt.Println("Don't have endpoint.")
	return ElasticResPost{}, errors.New("Don't have endpoint.\n")
}

func (e *elasticIndex) UpdateWithID(Type, ID string, Data interface{}) (ElasticResPost, error) {

	if e.EndPoint != "" {
		byteData, _ := json.Marshal(Data)
		postBody := bytes.NewReader(byteData)
		esHttp := elasticHttp{
			EndPoint: e.EndPoint,
			Index:    e.Index,
			Type:     Type,
			ID:       ID,
			Path:     "/" + e.Index + "/" + Type + "/" + ID,
		}

		res, err := esHttp.PUT(postBody)
		resBody, _ := ioutil.ReadAll(res.Body)
		var resData ElasticResPost
		json.Unmarshal(resBody, &resData)

		if resData.Result != "updated" && resData.Result != "created" {
			return resData, errors.New("Update failed.")
		}
		return resData, err
	}
	fmt.Println("Don't have endpoint.")
	return ElasticResPost{}, errors.New("Don't have endpoint.")
}

func (e *elasticIndex) SearchAll(Type string) (ElasticResGet, error) {
	if e.EndPoint != "" {
		esHttp := elasticHttp{
			EndPoint: e.EndPoint,
			Index:    e.Index,
			Type:     Type,
			ID:       "",
			Path:     "/" + e.Index + "/" + Type + "/_search",
		}

		res, err := esHttp.GET()
		resBody, _ := ioutil.ReadAll(res.Body)
		var resData ElasticResGet
		json.Unmarshal(resBody, &resData)
		return resData, err
	}
	fmt.Println("Don't have endpoint.")
	return ElasticResGet{}, errors.New("Don't have endpoint.\n")
}

func (e *ElasticClient) UnmarshalOnlyData(data ElasticResGet, out interface{}) {
	var datas []interface{}
	for _, temp := range data.Hits.Hits {
		datas = append(datas, temp.Source)
	}
	byteTemp, _ := json.Marshal(datas)
	json.Unmarshal(byteTemp, &out)
}

func (e *elasticIndex) Query(Type string, body io.Reader) (interface{}, int, error) {
	if e.EndPoint != "" {
		esHttp := elasticHttp{
			EndPoint: e.EndPoint,
			Index:    e.Index,
			Type:     Type,
			ID:       "",
			Path:     "/" + e.Index + "/" + Type + "/_search",
		}

		res, err := esHttp.Request(http.MethodGet, e.EndPoint+esHttp.Path, typeJSON, body)

		if err != nil {
			return nil, 500, err
		}

		resBody, _ := ioutil.ReadAll(res.Body)
		var resData interface{}
		json.Unmarshal(resBody, &resData)
		return resData, res.StatusCode, err
	}
	fmt.Println("Don't have endpoint.")
	return ElasticResGet{}, 500, errors.New("Don't have endpoint.\n")
}
