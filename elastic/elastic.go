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

// ElasticClient after use fun elastic.New(url)
type ElasticClient struct {
	EndPoint string
}

// New ElasticClient
func New(url string) ElasticClient {
	return ElasticClient{
		EndPoint: url,
	}
}

type elasticIndex struct {
	EndPoint string
	Index    string
}

// NewIndex func for NewClient With SetIndex
func (e *ElasticClient) NewIndex(Index string) ElasticIndex {
	if e.EndPoint == "" {
		fmt.Println("Don't have endpoint.")
		return ElasticIndex{}
	}
	return ElasticIndex{
		elasticIndex{
			EndPoint: e.EndPoint,
			Index:    Index,
		},
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
		esHTTP := elasticHTTP{
			EndPoint: e.EndPoint,
			Index:    e.Index,
			Type:     Type,
			ID:       ID,
			Path:     "/" + e.Index + "/" + Type + "/" + ID,
		}

		res, err := esHTTP.POST(postBody)
		resBody, _ := ioutil.ReadAll(res.Body)
		var resData ElasticResPost
		json.Unmarshal(resBody, &resData)

		if !resData.Created {
			return ElasticResPost{}, errors.New("insert failed")
		}
		return resData, err
	}
	fmt.Println("Don't have endpoint.")
	return ElasticResPost{}, errors.New("don't have endpoint")
}

// UpdateWithID Deprecated, Please use UpdateSomeDataWithID
func (e *elasticIndex) UpdateWithID(Type, ID string, Data interface{}) (ElasticResPost, error) {

	if e.EndPoint != "" {
		byteData, _ := json.Marshal(Data)
		postBody := bytes.NewReader(byteData)
		esHTTP := elasticHTTP{
			EndPoint: e.EndPoint,
			Index:    e.Index,
			Type:     Type,
			ID:       ID,
			Path:     "/" + e.Index + "/" + Type + "/" + ID,
		}

		res, err := esHTTP.PUT(postBody)
		resBody, _ := ioutil.ReadAll(res.Body)
		var resData ElasticResPost
		json.Unmarshal(resBody, &resData)

		if resData.Result != "updated" && resData.Result != "created" {
			return resData, errors.New("update failed")
		}
		return resData, err
	}
	fmt.Println("Don't have endpoint.")
	return ElasticResPost{}, errors.New("don't have endpoint")
}

func (e *elasticIndex) UpdateSomeDataWithID(Type, ID string, Data interface{}) (ElasticResPost, error) {

	if e.EndPoint != "" {
		type bodyUpdate struct {
			Doc interface{} `json:"doc"`
		}
		byteData, _ := json.Marshal(bodyUpdate{Doc: Data})
		postBody := bytes.NewReader(byteData)
		esHTTP := elasticHTTP{
			EndPoint: e.EndPoint,
			Index:    e.Index,
			Type:     Type,
			ID:       ID,
			Path:     "/" + e.Index + "/" + Type + "/" + ID + "/_update",
		}
		res, err := esHTTP.POST(postBody)
		resBody, _ := ioutil.ReadAll(res.Body)
		var resData ElasticResPost
		json.Unmarshal(resBody, &resData)

		if resData.Result != "updated" && resData.Result != "created" {
			return resData, errors.New("update failed")
		}
		return resData, err
	}
	fmt.Println("Don't have endpoint.")
	return ElasticResPost{}, errors.New("don't have endpoint")
}

func (e *elasticIndex) SearchAll(Type string) (ElasticResGet, error) {
	if e.EndPoint != "" {
		esHTTP := elasticHTTP{
			EndPoint: e.EndPoint,
			Index:    e.Index,
			Type:     Type,
			ID:       "",
			Path:     "/" + e.Index + "/" + Type + "/_search",
		}

		res, err := esHTTP.GET()
		resBody, _ := ioutil.ReadAll(res.Body)
		var resData ElasticResGet
		json.Unmarshal(resBody, &resData)
		return resData, err
	}
	fmt.Println("Don't have endpoint.")
	return ElasticResGet{}, errors.New("don't have endpoint")
}

// UnmarshalOnlyData is func mapping data in source
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
		esHTTP := elasticHTTP{
			EndPoint: e.EndPoint,
			Index:    e.Index,
			Type:     Type,
			ID:       "",
			Path:     "/" + e.Index + "/" + Type + "/_search",
		}

		res, err := esHTTP.Request(http.MethodGet, e.EndPoint+esHTTP.Path, typeJSON, body)

		if err != nil {
			return nil, 500, err
		}

		resBody, _ := ioutil.ReadAll(res.Body)
		var resData interface{}
		json.Unmarshal(resBody, &resData)
		return resData, res.StatusCode, err
	}
	fmt.Println("Don't have endpoint.")
	return ElasticResGet{}, 500, errors.New("don't have endpoint")
}

func (e *elasticIndex) GetByID(Type, ID string) (interface{}, int, error) {
	if e.EndPoint != "" {
		esHTTP := elasticHTTP{
			EndPoint: e.EndPoint,
			Index:    e.Index,
			Type:     Type,
			ID:       ID,
			Path:     "/" + e.Index + "/" + Type + "/" + ID,
		}

		res, err := esHTTP.Request(http.MethodGet, e.EndPoint+esHTTP.Path, typeJSON, nil)

		if err != nil {
			return nil, 500, err
		}

		resBody, _ := ioutil.ReadAll(res.Body)
		var resData interface{}
		json.Unmarshal(resBody, &resData)
		return resData, res.StatusCode, err
	}
	fmt.Println("Don't have endpoint.")
	return ElasticResGet{}, 500, errors.New("don't have endpoint")
}
