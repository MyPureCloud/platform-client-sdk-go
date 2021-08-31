package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkfetchrelationshipsresponse
type Bulkfetchrelationshipsresponse struct { 
	// Results
	Results *[]Bulkresponseresultrelationshipentity `json:"results,omitempty"`


	// ErrorCount
	ErrorCount *int `json:"errorCount,omitempty"`


	// ErrorIndexes
	ErrorIndexes *[]int `json:"errorIndexes,omitempty"`

}

func (o *Bulkfetchrelationshipsresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkfetchrelationshipsresponse
	
	return json.Marshal(&struct { 
		Results *[]Bulkresponseresultrelationshipentity `json:"results,omitempty"`
		
		ErrorCount *int `json:"errorCount,omitempty"`
		
		ErrorIndexes *[]int `json:"errorIndexes,omitempty"`
		*Alias
	}{ 
		Results: o.Results,
		
		ErrorCount: o.ErrorCount,
		
		ErrorIndexes: o.ErrorIndexes,
		Alias:    (*Alias)(o),
	})
}

func (o *Bulkfetchrelationshipsresponse) UnmarshalJSON(b []byte) error {
	var BulkfetchrelationshipsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BulkfetchrelationshipsresponseMap)
	if err != nil {
		return err
	}
	
	if Results, ok := BulkfetchrelationshipsresponseMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	
	if ErrorCount, ok := BulkfetchrelationshipsresponseMap["errorCount"].(float64); ok {
		ErrorCountInt := int(ErrorCount)
		o.ErrorCount = &ErrorCountInt
	}
	
	if ErrorIndexes, ok := BulkfetchrelationshipsresponseMap["errorIndexes"].([]interface{}); ok {
		ErrorIndexesString, _ := json.Marshal(ErrorIndexes)
		json.Unmarshal(ErrorIndexesString, &o.ErrorIndexes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkfetchrelationshipsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
