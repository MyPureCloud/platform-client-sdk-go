package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Statisticalsummary
type Statisticalsummary struct { 
	// Max
	Max *float32 `json:"max,omitempty"`


	// Min
	Min *float32 `json:"min,omitempty"`


	// Count
	Count *int `json:"count,omitempty"`


	// CountNegative
	CountNegative *int `json:"countNegative,omitempty"`


	// CountPositive
	CountPositive *int `json:"countPositive,omitempty"`


	// Sum
	Sum *float32 `json:"sum,omitempty"`


	// Current
	Current *float32 `json:"current,omitempty"`


	// Ratio
	Ratio *float32 `json:"ratio,omitempty"`


	// Numerator
	Numerator *float32 `json:"numerator,omitempty"`


	// Denominator
	Denominator *float32 `json:"denominator,omitempty"`


	// Target
	Target *float32 `json:"target,omitempty"`

}

func (o *Statisticalsummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Statisticalsummary
	
	return json.Marshal(&struct { 
		Max *float32 `json:"max,omitempty"`
		
		Min *float32 `json:"min,omitempty"`
		
		Count *int `json:"count,omitempty"`
		
		CountNegative *int `json:"countNegative,omitempty"`
		
		CountPositive *int `json:"countPositive,omitempty"`
		
		Sum *float32 `json:"sum,omitempty"`
		
		Current *float32 `json:"current,omitempty"`
		
		Ratio *float32 `json:"ratio,omitempty"`
		
		Numerator *float32 `json:"numerator,omitempty"`
		
		Denominator *float32 `json:"denominator,omitempty"`
		
		Target *float32 `json:"target,omitempty"`
		*Alias
	}{ 
		Max: o.Max,
		
		Min: o.Min,
		
		Count: o.Count,
		
		CountNegative: o.CountNegative,
		
		CountPositive: o.CountPositive,
		
		Sum: o.Sum,
		
		Current: o.Current,
		
		Ratio: o.Ratio,
		
		Numerator: o.Numerator,
		
		Denominator: o.Denominator,
		
		Target: o.Target,
		Alias:    (*Alias)(o),
	})
}

func (o *Statisticalsummary) UnmarshalJSON(b []byte) error {
	var StatisticalsummaryMap map[string]interface{}
	err := json.Unmarshal(b, &StatisticalsummaryMap)
	if err != nil {
		return err
	}
	
	if Max, ok := StatisticalsummaryMap["max"].(float64); ok {
		MaxFloat32 := float32(Max)
		o.Max = &MaxFloat32
	}
    
	if Min, ok := StatisticalsummaryMap["min"].(float64); ok {
		MinFloat32 := float32(Min)
		o.Min = &MinFloat32
	}
    
	if Count, ok := StatisticalsummaryMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	
	if CountNegative, ok := StatisticalsummaryMap["countNegative"].(float64); ok {
		CountNegativeInt := int(CountNegative)
		o.CountNegative = &CountNegativeInt
	}
	
	if CountPositive, ok := StatisticalsummaryMap["countPositive"].(float64); ok {
		CountPositiveInt := int(CountPositive)
		o.CountPositive = &CountPositiveInt
	}
	
	if Sum, ok := StatisticalsummaryMap["sum"].(float64); ok {
		SumFloat32 := float32(Sum)
		o.Sum = &SumFloat32
	}
    
	if Current, ok := StatisticalsummaryMap["current"].(float64); ok {
		CurrentFloat32 := float32(Current)
		o.Current = &CurrentFloat32
	}
    
	if Ratio, ok := StatisticalsummaryMap["ratio"].(float64); ok {
		RatioFloat32 := float32(Ratio)
		o.Ratio = &RatioFloat32
	}
    
	if Numerator, ok := StatisticalsummaryMap["numerator"].(float64); ok {
		NumeratorFloat32 := float32(Numerator)
		o.Numerator = &NumeratorFloat32
	}
    
	if Denominator, ok := StatisticalsummaryMap["denominator"].(float64); ok {
		DenominatorFloat32 := float32(Denominator)
		o.Denominator = &DenominatorFloat32
	}
    
	if Target, ok := StatisticalsummaryMap["target"].(float64); ok {
		TargetFloat32 := float32(Target)
		o.Target = &TargetFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Statisticalsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
