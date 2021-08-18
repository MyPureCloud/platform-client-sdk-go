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

func (u *Statisticalsummary) MarshalJSON() ([]byte, error) {
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
		Max: u.Max,
		
		Min: u.Min,
		
		Count: u.Count,
		
		CountNegative: u.CountNegative,
		
		CountPositive: u.CountPositive,
		
		Sum: u.Sum,
		
		Current: u.Current,
		
		Ratio: u.Ratio,
		
		Numerator: u.Numerator,
		
		Denominator: u.Denominator,
		
		Target: u.Target,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Statisticalsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
