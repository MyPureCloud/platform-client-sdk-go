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
	Min *int `json:"min,omitempty"`


	// Max - The maximum value in this metric
	Max *int `json:"max,omitempty"`


	// Sum - The total of the values for this metric
	Sum *int `json:"sum,omitempty"`

}

func (o *Developmentactivityaggregatequeryresponsestatistics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Developmentactivityaggregatequeryresponsestatistics
	
	return json.Marshal(&struct { 
		Count *int `json:"count,omitempty"`
		
		Min *int `json:"min,omitempty"`
		
		Max *int `json:"max,omitempty"`
		
		Sum *int `json:"sum,omitempty"`
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
		MinInt := int(Min)
		o.Min = &MinInt
	}
	
	if Max, ok := DevelopmentactivityaggregatequeryresponsestatisticsMap["max"].(float64); ok {
		MaxInt := int(Max)
		o.Max = &MaxInt
	}
	
	if Sum, ok := DevelopmentactivityaggregatequeryresponsestatisticsMap["sum"].(float64); ok {
		SumInt := int(Sum)
		o.Sum = &SumInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryresponsestatistics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
