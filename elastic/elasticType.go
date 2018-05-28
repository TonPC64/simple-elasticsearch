package elastic

type ElasticResGet struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits         *ElasticHit  `json:"hits"`
	Aggregations *ElasticAggr `json:"aggregations"`
}

type ElasticHit struct {
	Total    int `json:"total"`
	MaxScore int `json:"max_score"`
	Hits     []struct {
		Index  string      `json:"_index"`
		Type   string      `json:"_type"`
		ID     string      `json:"_id"`
		Score  int         `json:"_score"`
		Source interface{} `json:"_source"`
	} `json:"hits"`
}

type ElasticAggr struct {
	ByType struct {
		DocCount  int `json:"doc_count"`
		ByAdminID struct {
			DocCount   int `json:"doc_count"`
			TimeSeries struct {
				DocCount      int `json:"doc_count"`
				ActivePerHour struct {
					Buckets []struct {
						KeyAsString string `json:"key_as_string"`
						Key         int64  `json:"key"`
						DocCount    int    `json:"doc_count"`
					} `json:"buckets"`
				} `json:"active_per_hour"`
			} `json:"time_series"`
		} `json:"byAdmin_id"`
	} `json:"byType"`
}

type ElasticResPost struct {
	ID     string `json:"_id"`
	Index  string `json:"_index"`
	Shards struct {
		Failed     int `json:"failed"`
		Successful int `json:"successful"`
		Total      int `json:"total"`
	} `json:"_shards"`
	Type    string `json:"_type"`
	Version int    `json:"_version"`
	Created bool   `json:"created"`
	Result  string `json:"result"`
}
