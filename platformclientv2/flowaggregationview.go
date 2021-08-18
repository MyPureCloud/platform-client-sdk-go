package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowaggregationview
type Flowaggregationview struct { 
	// Target - Target metric name
	Target *string `json:"target,omitempty"`


	// Name - A unique name for this view. Must be distinct from other views and built-in metric names.
	Name *string `json:"name,omitempty"`


	// Function - Type of view you wish to create
	Function *string `json:"function,omitempty"`


	// VarRange - Range of numbers for slicing up data
	VarRange *Aggregationrange `json:"range,omitempty"`

}

func (u *Flowaggregationview) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowaggregationview

	

	return json.Marshal(&struct { 
		Target *string `json:"target,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Function *string `json:"function,omitempty"`
		
		VarRange *Aggregationrange `json:"range,omitempty"`
		*Alias
	}{ 
		Target: u.Target,
		
		Name: u.Name,
		
		Function: u.Function,
		
		VarRange: u.VarRange,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Flowaggregationview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
