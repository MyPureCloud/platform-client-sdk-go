package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkcallbackpatchresponse
type Bulkcallbackpatchresponse struct { 
	// Results - A list of the results from the bulk operation.
	Results *[]Bulkresult `json:"results,omitempty"`


	// ErrorCount - The number of errors from the bulk operation.
	ErrorCount *int `json:"errorCount,omitempty"`


	// ErrorIndexes - An index of where the errors are in the listing.
	ErrorIndexes *[]int `json:"errorIndexes,omitempty"`

}

func (o *Bulkcallbackpatchresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkcallbackpatchresponse
	
	return json.Marshal(&struct { 
		Results *[]Bulkresult `json:"results,omitempty"`
		
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

func (o *Bulkcallbackpatchresponse) UnmarshalJSON(b []byte) error {
	var BulkcallbackpatchresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BulkcallbackpatchresponseMap)
	if err != nil {
		return err
	}
	
	if Results, ok := BulkcallbackpatchresponseMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	
	if ErrorCount, ok := BulkcallbackpatchresponseMap["errorCount"].(float64); ok {
		ErrorCountInt := int(ErrorCount)
		o.ErrorCount = &ErrorCountInt
	}
	
	if ErrorIndexes, ok := BulkcallbackpatchresponseMap["errorIndexes"].([]interface{}); ok {
		ErrorIndexesString, _ := json.Marshal(ErrorIndexes)
		json.Unmarshal(ErrorIndexesString, &o.ErrorIndexes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkcallbackpatchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
