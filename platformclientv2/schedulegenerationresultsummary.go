package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulegenerationresultsummary
type Schedulegenerationresultsummary struct { 
	// Failed - Whether the schedule generation run failed
	Failed *bool `json:"failed,omitempty"`


	// RunId - The run ID for the schedule generation. Reference this when requesting support
	RunId *string `json:"runId,omitempty"`


	// MessageCount - The number of schedule generation messages for this schedule generation run
	MessageCount *int `json:"messageCount,omitempty"`

}

func (u *Schedulegenerationresultsummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Schedulegenerationresultsummary

	

	return json.Marshal(&struct { 
		Failed *bool `json:"failed,omitempty"`
		
		RunId *string `json:"runId,omitempty"`
		
		MessageCount *int `json:"messageCount,omitempty"`
		*Alias
	}{ 
		Failed: u.Failed,
		
		RunId: u.RunId,
		
		MessageCount: u.MessageCount,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Schedulegenerationresultsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
