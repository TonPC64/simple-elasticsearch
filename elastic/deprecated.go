package elastic

import (
	"errors"
	"fmt"
)

// SetIndex Deprecated function (Use NewIndex)
func (e *ElasticClient) SetIndex(Index string) (elasticIndex, error) {
	if e.EndPoint == "" {
		fmt.Println("Don't have endpoint.")
		return elasticIndex{}, errors.New("don't have endpoint")
	}
	return elasticIndex{
		EndPoint: e.EndPoint,
		Index:    Index,
	}, nil
}
