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
	Min *float32 `json:"min,omitempty"`


	// Max - The maximum value in this metric
	Max *float32 `json:"max,omitempty"`


	// Sum - The total of the values for this metric
	Sum *float32 `json:"sum,omitempty"`

}

func (o *Learningassignmentaggregatequeryresponsestats) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentaggregatequeryresponsestats
	
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
		MinFloat32 := float32(Min)
		o.Min = &MinFloat32
	}
	
	if Max, ok := LearningassignmentaggregatequeryresponsestatsMap["max"].(float64); ok {
		MaxFloat32 := float32(Max)
		o.Max = &MaxFloat32
	}
	
	if Sum, ok := LearningassignmentaggregatequeryresponsestatsMap["sum"].(float64); ok {
		SumFloat32 := float32(Sum)
		o.Sum = &SumFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregatequeryresponsestats) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
