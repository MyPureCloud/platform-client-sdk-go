package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Durationcondition
type Durationcondition struct { 
	// DurationTarget
	DurationTarget *string `json:"durationTarget,omitempty"`


	// DurationOperator
	DurationOperator *string `json:"durationOperator,omitempty"`


	// DurationRange
	DurationRange *string `json:"durationRange,omitempty"`

}

func (u *Durationcondition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Durationcondition

	

	return json.Marshal(&struct { 
		DurationTarget *string `json:"durationTarget,omitempty"`
		
		DurationOperator *string `json:"durationOperator,omitempty"`
		
		DurationRange *string `json:"durationRange,omitempty"`
		*Alias
	}{ 
		DurationTarget: u.DurationTarget,
		
		DurationOperator: u.DurationOperator,
		
		DurationRange: u.DurationRange,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Durationcondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
