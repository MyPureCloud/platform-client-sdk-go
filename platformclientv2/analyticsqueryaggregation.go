package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsqueryaggregation
type Analyticsqueryaggregation struct { 
	// VarType - Optional type, can usually be inferred
	VarType *string `json:"type,omitempty"`


	// Dimension - For use with termFrequency aggregations
	Dimension *string `json:"dimension,omitempty"`


	// Metric - For use with numericRange aggregations
	Metric *string `json:"metric,omitempty"`


	// Size - For use with termFrequency aggregations
	Size *int `json:"size,omitempty"`


	// Ranges - For use with numericRange aggregations
	Ranges *[]Aggregationrange `json:"ranges,omitempty"`

}

func (o *Analyticsqueryaggregation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsqueryaggregation
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Dimension *string `json:"dimension,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		Size *int `json:"size,omitempty"`
		
		Ranges *[]Aggregationrange `json:"ranges,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Dimension: o.Dimension,
		
		Metric: o.Metric,
		
		Size: o.Size,
		
		Ranges: o.Ranges,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticsqueryaggregation) UnmarshalJSON(b []byte) error {
	var AnalyticsqueryaggregationMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsqueryaggregationMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := AnalyticsqueryaggregationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Dimension, ok := AnalyticsqueryaggregationMap["dimension"].(string); ok {
		o.Dimension = &Dimension
	}
    
	if Metric, ok := AnalyticsqueryaggregationMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if Size, ok := AnalyticsqueryaggregationMap["size"].(float64); ok {
		SizeInt := int(Size)
		o.Size = &SizeInt
	}
	
	if Ranges, ok := AnalyticsqueryaggregationMap["ranges"].([]interface{}); ok {
		RangesString, _ := json.Marshal(Ranges)
		json.Unmarshal(RangesString, &o.Ranges)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsqueryaggregation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
