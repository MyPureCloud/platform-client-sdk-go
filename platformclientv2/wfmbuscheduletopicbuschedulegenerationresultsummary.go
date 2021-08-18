package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuscheduletopicbuschedulegenerationresultsummary
type Wfmbuscheduletopicbuschedulegenerationresultsummary struct { 
	// Failed
	Failed *bool `json:"failed,omitempty"`


	// RunId
	RunId *string `json:"runId,omitempty"`


	// MessageCount
	MessageCount *int `json:"messageCount,omitempty"`

}

func (u *Wfmbuscheduletopicbuschedulegenerationresultsummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuscheduletopicbuschedulegenerationresultsummary

	

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
func (o *Wfmbuscheduletopicbuschedulegenerationresultsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
