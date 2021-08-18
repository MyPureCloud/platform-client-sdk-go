package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulingtestingoptionsrequest
type Schedulingtestingoptionsrequest struct { 
	// FastScheduling - Whether to enable fast scheduling
	FastScheduling *bool `json:"fastScheduling,omitempty"`


	// DelayScheduling - Whether to force delayed scheduling
	DelayScheduling *bool `json:"delayScheduling,omitempty"`


	// FailScheduling - Whether to force scheduling to fail
	FailScheduling *bool `json:"failScheduling,omitempty"`


	// PopulateWarnings - Whether to populate warnings in the generated schedule
	PopulateWarnings *bool `json:"populateWarnings,omitempty"`

}

func (u *Schedulingtestingoptionsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Schedulingtestingoptionsrequest

	

	return json.Marshal(&struct { 
		FastScheduling *bool `json:"fastScheduling,omitempty"`
		
		DelayScheduling *bool `json:"delayScheduling,omitempty"`
		
		FailScheduling *bool `json:"failScheduling,omitempty"`
		
		PopulateWarnings *bool `json:"populateWarnings,omitempty"`
		*Alias
	}{ 
		FastScheduling: u.FastScheduling,
		
		DelayScheduling: u.DelayScheduling,
		
		FailScheduling: u.FailScheduling,
		
		PopulateWarnings: u.PopulateWarnings,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Schedulingtestingoptionsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
