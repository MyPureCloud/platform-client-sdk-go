package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkcontactsresponse
type Bulkcontactsresponse struct { 
	// Results
	Results *[]Bulkresponseresultexternalcontactexternalcontact `json:"results,omitempty"`


	// ErrorCount
	ErrorCount *int `json:"errorCount,omitempty"`


	// ErrorIndexes
	ErrorIndexes *[]int `json:"errorIndexes,omitempty"`

}

func (o *Bulkcontactsresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkcontactsresponse
	
	return json.Marshal(&struct { 
		Results *[]Bulkresponseresultexternalcontactexternalcontact `json:"results,omitempty"`
		
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

func (o *Bulkcontactsresponse) UnmarshalJSON(b []byte) error {
	var BulkcontactsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BulkcontactsresponseMap)
	if err != nil {
		return err
	}
	
	if Results, ok := BulkcontactsresponseMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	
	if ErrorCount, ok := BulkcontactsresponseMap["errorCount"].(float64); ok {
		ErrorCountInt := int(ErrorCount)
		o.ErrorCount = &ErrorCountInt
	}
	
	if ErrorIndexes, ok := BulkcontactsresponseMap["errorIndexes"].([]interface{}); ok {
		ErrorIndexesString, _ := json.Marshal(ErrorIndexes)
		json.Unmarshal(ErrorIndexesString, &o.ErrorIndexes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkcontactsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
