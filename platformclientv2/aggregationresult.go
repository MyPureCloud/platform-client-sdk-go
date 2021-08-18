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

func (u *Aggregationresult) MarshalJSON() ([]byte, error) {
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
		VarType: u.VarType,
		
		Dimension: u.Dimension,
		
		Metric: u.Metric,
		
		Count: u.Count,
		
		Results: u.Results,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Aggregationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
