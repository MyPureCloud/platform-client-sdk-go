package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Developmentactivityaggregatequeryresponsestatistics
type Developmentactivityaggregatequeryresponsestatistics struct { 
	// Count - The count for this metric
	Count *int `json:"count,omitempty"`


	// Min - The minimum value in this metric
	Min *float32 `json:"min,omitempty"`


	// Max - The maximum value in this metric
	Max *float32 `json:"max,omitempty"`


	// Sum - The total of the values for this metric
	Sum *float32 `json:"sum,omitempty"`

}

func (o *Developmentactivityaggregatequeryresponsestatistics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Developmentactivityaggregatequeryresponsestatistics
	
	return json.Marshal(&struct { 
		Count *int `json:"count,omitempty"`
		
		Min *float32 `json:"min,omitempty"`
		
		Max *float32 `json:"max,omitempty"`
		
		Sum *float32 `json:"sum,omitempty"`
		*Alias
	}{ 
		Count: o.Count,
		
		Min: o.Min,
		
		Max: o.Max,
		
		Sum: o.Sum,
		Alias:    (*Alias)(o),
	})
}

func (o *Developmentactivityaggregatequeryresponsestatistics) UnmarshalJSON(b []byte) error {
	var DevelopmentactivityaggregatequeryresponsestatisticsMap map[string]interface{}
	err := json.Unmarshal(b, &DevelopmentactivityaggregatequeryresponsestatisticsMap)
	if err != nil {
		return err
	}
	
	if Count, ok := DevelopmentactivityaggregatequeryresponsestatisticsMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	
	if Min, ok := DevelopmentactivityaggregatequeryresponsestatisticsMap["min"].(float64); ok {
		MinFloat32 := float32(Min)
		o.Min = &MinFloat32
	}
	
	if Max, ok := DevelopmentactivityaggregatequeryresponsestatisticsMap["max"].(float64); ok {
		MaxFloat32 := float32(Max)
		o.Max = &MaxFloat32
	}
	
	if Sum, ok := DevelopmentactivityaggregatequeryresponsestatisticsMap["sum"].(float64); ok {
		SumFloat32 := float32(Sum)
		o.Sum = &SumFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryresponsestatistics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
