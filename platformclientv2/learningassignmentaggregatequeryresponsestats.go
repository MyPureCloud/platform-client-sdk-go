package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentaggregatequeryresponsestats
type Learningassignmentaggregatequeryresponsestats struct { 
	// Count - The count for this metric
	Count *int `json:"count,omitempty"`


	// Min - The minimum value in this metric
	Min *int `json:"min,omitempty"`


	// Max - The maximum value in this metric
	Max *int `json:"max,omitempty"`


	// Sum - The total of the values for this metric
	Sum *int `json:"sum,omitempty"`

}

func (o *Learningassignmentaggregatequeryresponsestats) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentaggregatequeryresponsestats
	
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

func (o *Learningassignmentaggregatequeryresponsestats) UnmarshalJSON(b []byte) error {
	var LearningassignmentaggregatequeryresponsestatsMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmentaggregatequeryresponsestatsMap)
	if err != nil {
		return err
	}
	
	if Count, ok := LearningassignmentaggregatequeryresponsestatsMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	
	if Min, ok := LearningassignmentaggregatequeryresponsestatsMap["min"].(float64); ok {
		MinInt := int(Min)
		o.Min = &MinInt
	}
	
	if Max, ok := LearningassignmentaggregatequeryresponsestatsMap["max"].(float64); ok {
		MaxInt := int(Max)
		o.Max = &MaxInt
	}
	
	if Sum, ok := LearningassignmentaggregatequeryresponsestatsMap["sum"].(float64); ok {
		SumInt := int(Sum)
		o.Sum = &SumInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregatequeryresponsestats) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
