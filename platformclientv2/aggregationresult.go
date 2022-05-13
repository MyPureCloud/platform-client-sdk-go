package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Aggregationresult
type Aggregationresult struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Dimension - For termFrequency aggregations
	Dimension *string `json:"dimension,omitempty"`


	// Metric - For numericRange aggregations
	Metric *string `json:"metric,omitempty"`


	// Count
	Count *int `json:"count,omitempty"`


	// Results
	Results *[]Aggregationresultentry `json:"results,omitempty"`

}

func (o *Aggregationresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Aggregationresult
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Dimension *string `json:"dimension,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		Count *int `json:"count,omitempty"`
		
		Results *[]Aggregationresultentry `json:"results,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Dimension: o.Dimension,
		
		Metric: o.Metric,
		
		Count: o.Count,
		
		Results: o.Results,
		Alias:    (*Alias)(o),
	})
}

func (o *Aggregationresult) UnmarshalJSON(b []byte) error {
	var AggregationresultMap map[string]interface{}
	err := json.Unmarshal(b, &AggregationresultMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := AggregationresultMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Dimension, ok := AggregationresultMap["dimension"].(string); ok {
		o.Dimension = &Dimension
	}
    
	if Metric, ok := AggregationresultMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if Count, ok := AggregationresultMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	
	if Results, ok := AggregationresultMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Aggregationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
