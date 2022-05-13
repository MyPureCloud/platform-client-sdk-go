package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Aggregationresultentry
type Aggregationresultentry struct { 
	// Count
	Count *int `json:"count,omitempty"`


	// Value - For termFrequency aggregations
	Value *string `json:"value,omitempty"`


	// Gte - For numericRange aggregations
	Gte *float32 `json:"gte,omitempty"`


	// Lt - For numericRange aggregations
	Lt *float32 `json:"lt,omitempty"`

}

func (o *Aggregationresultentry) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Aggregationresultentry
	
	return json.Marshal(&struct { 
		Count *int `json:"count,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		Gte *float32 `json:"gte,omitempty"`
		
		Lt *float32 `json:"lt,omitempty"`
		*Alias
	}{ 
		Count: o.Count,
		
		Value: o.Value,
		
		Gte: o.Gte,
		
		Lt: o.Lt,
		Alias:    (*Alias)(o),
	})
}

func (o *Aggregationresultentry) UnmarshalJSON(b []byte) error {
	var AggregationresultentryMap map[string]interface{}
	err := json.Unmarshal(b, &AggregationresultentryMap)
	if err != nil {
		return err
	}
	
	if Count, ok := AggregationresultentryMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	
	if Value, ok := AggregationresultentryMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if Gte, ok := AggregationresultentryMap["gte"].(float64); ok {
		GteFloat32 := float32(Gte)
		o.Gte = &GteFloat32
	}
    
	if Lt, ok := AggregationresultentryMap["lt"].(float64); ok {
		LtFloat32 := float32(Lt)
		o.Lt = &LtFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Aggregationresultentry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
